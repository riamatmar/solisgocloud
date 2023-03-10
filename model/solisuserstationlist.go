package model

type SolisUserStationList struct {
	Success bool   `json:"success"`
	Code    string `json:"code"`
	Msg     string `json:"msg"`
	Data    Data   `json:"data"`
}

type Data struct {
	StationStatusVo StationStatusVo `json:"stationStatusVo"`
	Page            Page            `json:"page"`
	MpptSwitch      int64           `json:"mpptSwitch"`
}

type Page struct {
	Records          []Record      `json:"records"`
	Total            int64         `json:"total"`
	Size             int64         `json:"size"`
	Current          int64         `json:"current"`
	Orders           []interface{} `json:"orders"`
	OptimizeCountSQL bool          `json:"optimizeCountSql"`
	SearchCount      bool          `json:"searchCount"`
	Pages            int64         `json:"pages"`
}

type Record struct {
	ID                          string  `json:"id"`
	DataTimestamp               string  `json:"dataTimestamp"`
	DataTimestampStr            string  `json:"dataTimestampStr"`
	FullHour                    float64 `json:"fullHour"`
	DayPowerGeneration          float64 `json:"dayPowerGeneration"`
	MonthCarbonDioxide          float64 `json:"monthCarbonDioxide"`
	UserID                      string  `json:"userId"`
	Sno                         string  `json:"sno"`
	CountryStr                  string  `json:"countryStr"`
	RegionStr                   string  `json:"regionStr"`
	CityStr                     string  `json:"cityStr"`
	CountyStr                   string  `json:"countyStr"`
	State                       int64   `json:"state"`
	DIP                         float64 `json:"dip"`
	Azimuth                     float64 `json:"azimuth"`
	Power                       float64 `json:"power"`
	TimeZone                    float64 `json:"timeZone"`
	TimeZoneName                string  `json:"timeZoneName"`
	TimeZoneStr                 string  `json:"timeZoneStr"`
	TimeZoneID                  string  `json:"timeZoneId"`
	Daylight                    float64 `json:"daylight"`
	PowerStr                    string  `json:"powerStr"`
	CreateDate                  int64   `json:"createDate"`
	CreateDateStr               string  `json:"createDateStr"`
	Price                       float64 `json:"price"`
	Capacity                    float64 `json:"capacity"`
	CapacityStr                 string  `json:"capacityStr"`
	CapacityPercent             float64 `json:"capacityPercent"`
	Capacity1                   float64 `json:"capacity1"`
	PicName                     string  `json:"picName"`
	Pic1URL                     string  `json:"pic1Url"`
	DayEnergy                   float64 `json:"dayEnergy"`
	DayEnergyStr                string  `json:"dayEnergyStr"`
	DayIncome                   float64 `json:"dayIncome"`
	MonthEnergy                 float64 `json:"monthEnergy"`
	MonthEnergyStr              string  `json:"monthEnergyStr"`
	YearEnergy                  float64 `json:"yearEnergy"`
	YearEnergyStr               string  `json:"yearEnergyStr"`
	AllEnergy                   float64 `json:"allEnergy"`
	AllEnergyStr                string  `json:"allEnergyStr"`
	AllEnergy1                  float64 `json:"allEnergy1"`
	AllIncome                   float64 `json:"allIncome"`
	UpdateDate                  int64   `json:"updateDate"`
	Type                        int64   `json:"type"`
	SynchronizationType         int64   `json:"synchronizationType"`
	EpmType                     int64   `json:"epmType"`
	GridSwitch                  int64   `json:"gridSwitch"`
	GridSwitch1                 int64   `json:"gridSwitch1"`
	ShareProcess                int64   `json:"shareProcess"`
	AlarmLongStr                string  `json:"alarmLongStr"`
	DcInputType                 int64   `json:"dcInputType"`
	StationTypeNew              int64   `json:"stationTypeNew"`
	BatteryTotalDischargeEnergy float64 `json:"batteryTotalDischargeEnergy"`
	BatteryTotalChargeEnergy    float64 `json:"batteryTotalChargeEnergy"`
	GridPurchasedTotalEnergy    float64 `json:"gridPurchasedTotalEnergy"`
	GridSellTotalEnergy         float64 `json:"gridSellTotalEnergy"`
	HomeLoadTotalEnergy         float64 `json:"homeLoadTotalEnergy"`
	OneSelf                     float64 `json:"oneSelf"`
	BatteryTodayDischargeEnergy float64 `json:"batteryTodayDischargeEnergy"`
	BatteryTodayChargeEnergy    float64 `json:"batteryTodayChargeEnergy"`
	GridPurchasedTodayEnergy    float64 `json:"gridPurchasedTodayEnergy"`
	GridSellTodayEnergy         float64 `json:"gridSellTodayEnergy"`
	HomeLoadTodayEnergy         float64 `json:"homeLoadTodayEnergy"`
	OneSelfTotal                float64 `json:"oneSelfTotal"`
	Money                       string  `json:"money"`
	Remark1                     string  `json:"remark1"`
	Remark2                     string  `json:"remark2"`
	Remark3                     string  `json:"remark3"`
	SimFlowState                int64   `json:"simFlowState"`
	JxbType                     int64   `json:"jxbType"`
	InverterCount               int64   `json:"inverterCount"`
	Power1                      float64 `json:"power1"`
	MonthEnergy1                float64 `json:"monthEnergy1"`
	YearEnergy1                 float64 `json:"yearEnergy1"`
	DayEnergy1                  float64 `json:"dayEnergy1"`
}

type StationStatusVo struct {
	All      int64 `json:"all"`
	Normal   int64 `json:"normal"`
	Fault    int64 `json:"fault"`
	Offline  int64 `json:"offline"`
	Building int64 `json:"building"`
	Mppt     int64 `json:"mppt"`
}
