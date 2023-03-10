package solisgocloud

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/riamatmar/solisgocloud/model"
	"github.com/rs/zerolog/log"
)

type SolisCloudService interface {
	ReadUserStationListSolisCloud(pageNo int, pageSize int) (model.SolisUserStationList, error)
}

type solisCloudService struct {
	solisCloud model.SolisCloud
}

func (s *solisCloudService) ReadUserStationListSolisCloud(pageNo int, pageSize int) (model.SolisUserStationList, error) {

	solisUserStationList := model.SolisUserStationList{}

	paginator := model.SolisPaginator{
		PageNo:   pageNo,
		PageSize: pageSize,
	}

	paginatorJson, err := json.Marshal(paginator)
	if err != nil {
		log.Error().Err(err).Msg("ReadUserStationListSolisCloud")
		return solisUserStationList, err
	}

	contentMD5 := s.createContentMD5(string(paginatorJson))
	date := time.Now().UTC().Format(http.TimeFormat)
	path := "/v1/api/userStationList"

	param := fmt.Sprintf("POST\n%s\napplication/json\n%s\n%s", contentMD5, date, path)

	sign := s.calculateAuthorization(param)
	client := resty.New()

	resp, err := client.R().
		SetHeader("Authorization", "API "+s.solisCloud.KeyID+":"+sign).
		SetHeader("Content-MD5", contentMD5).
		SetHeader("Content-type", "application/json;charset=UTF-8").
		SetHeader("Date", date).
		SetBody(paginator).
		Post(s.solisCloud.URL + path)

	if err != nil {
		log.Error().Err(err).Msg("ReadUserStationListSolisCloud")
		return solisUserStationList, err
	}

	if resp.StatusCode() != http.StatusOK {
		log.Debug().Msg(string(resp.Body()))
		err = fmt.Errorf("error get read userstationlist (%d)", resp.StatusCode())
		log.Error().Err(err).Msg("ReadUserStationListSolisCloud")
		return solisUserStationList, err
	}

	err = json.Unmarshal(resp.Body(), &solisUserStationList)
	if err != nil {
		log.Debug().Msg(string(resp.Body()))
		log.Error().Err(err).Msg("ReadUserStationListSolisCloud")
		return model.SolisUserStationList{}, err
	}

	return solisUserStationList, nil
}

func (s *solisCloudService) createContentMD5(contentMD5 string) string {

	hasher := md5.New()
	hasher.Write([]byte(contentMD5))
	encodedString := base64.StdEncoding.EncodeToString(hasher.Sum(nil))

	return encodedString

}

func (s *solisCloudService) calculateAuthorization(input string) string {

	key_for_sign := []byte(s.solisCloud.KeySercret)
	h := hmac.New(sha1.New, key_for_sign)
	h.Write([]byte(input))
	authorization := base64.StdEncoding.EncodeToString(h.Sum(nil))

	return authorization

}

func CreateSolisCloudService(solisCloud model.SolisCloud) SolisCloudService {
	return &solisCloudService{
		solisCloud: solisCloud,
	}
}
