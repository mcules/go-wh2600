package models

type WeatherData struct {
	TempF          float64 `form:"tempf"`
	TempC          float64
	Humidity       float64 `form:"humidity"`
	DewptF         float64 `form:"dewptf"`
	DewptC         float64
	WindChillF     float64 `form:"windchillf"`
	WindChillC     float64
	WindDir        uint8   `form:"winddir"`
	WindSpeedMPH   float64 `form:"windspeedmph"`
	WindSpeedKmh   float64
	WindGustMPH    float64 `form:"windgustmph"`
	WindGustKmh    float64
	Rainin         float64 `form:"rainin"`
	DailyRainin    float64 `form:"dailyrainin"`
	WeeklyRainin   float64 `form:"weeklyrainin"`
	MonthlyRainin  float64 `form:"monthlyrainin"`
	YearlyRainin   float64 `form:"yearlyrainin"`
	SolarRadiation float64 `form:"solarradiation"`
	UV             uint8   `form:"UV"`
	IndoorTempF    float64 `form:"indoortempf"`
	IndoorTempC    float64
	IndoorHumidity uint8   `form:"indoorhumidity"`
	BaroMin        float64 `form:"baromin"`
	DateUTC        string  `form:"dateutc"`
	SoftwareType   string  `form:"softwaretype"`
	Action         string  `form:"action"`
	Realtime       bool    `form:"realtime"`
	RtFreq         uint8   `form:"rtfreq"`
}

func (wd *WeatherData) Recalc() {
	wd.TempC = FahrenheitToCelsius(wd.TempF)
	wd.DewptC = FahrenheitToCelsius(wd.DewptF)
	wd.IndoorTempC = FahrenheitToCelsius(wd.IndoorTempF)
	wd.WindChillC = FahrenheitToCelsius(wd.WindChillF)
	wd.WindSpeedKmh = MphToKmh(wd.WindSpeedMPH)
	wd.WindGustKmh = MphToKmh(wd.WindGustMPH)
}
