package models

type WeatherData struct {
	TempF          float32 `form:"tempf"`
	TempC          float32
	Humidity       float32 `form:"humidity"`
	DewptF         float32 `form:"dewptf"`
	DewptC         float32
	WindChillF     float32 `form:"windchillf"`
	WindChillC     float32
	WindDir        uint8   `form:"winddir"`
	WindSpeedMPH   float32 `form:"windspeedmph"`
	WindSpeedKmh   float32
	WindGustMPH    float32 `form:"windgustmph"`
	WindGustKmh    float32
	Rainin         float32 `form:"rainin"`
	DailyRainin    float32 `form:"dailyrainin"`
	WeeklyRainin   float32 `form:"weeklyrainin"`
	MonthlyRainin  float32 `form:"monthlyrainin"`
	YearlyRainin   float32 `form:"yearlyrainin"`
	SolarRadiation float32 `form:"solarradiation"`
	UV             uint8   `form:"UV"`
	IndoorTempF    float32 `form:"indoortempf"`
	IndoorTempC    float32
	IndoorHumidity uint8   `form:"indoorhumidity"`
	BaroMin        float32 `form:"baromin"`
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
