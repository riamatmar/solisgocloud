# solisgocloud
Golang package for SolisCloud API

###### Install
`go get github.com/riamatmar/solisgocloud`

## SolisCloud configuration:
This struct is needed for initialisation of the solisgocloud service.
``` golang
type SolisCloud struct {
	KeyID      string `json:"keyid"`
	KeySercret string `json:"keysercret"`
	URL        string `json:"url"`
}
```
## Initialisation of the solisgocloud service:
``` golang
solisCloudService := solisgocloud.CreateSolisCloudService(solisCloud)
if solisCloudService == nil {
  log.Fatal().Msg("unabbel to create solis cloud service")
}
```
## Read power station list (/v1/api/userStationList)
``` golang
solisUserStationList, err := solisCloudService.ReadUserStationListSolisCloud(1, 10)
if err != nil {
  log.Err(err).Msg("unabbel to read solisUserStationList")
}
```
