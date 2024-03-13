package constants

const (
	TIME_MILLISECONDS_OF_SECOND                              = 1000
	TIME_MILLISECONDS_OF_MINUTE                              = 60 * TIME_MILLISECONDS_OF_SECOND
	TIME_MILLISECONDS_OF_HOUR                                = 60 * TIME_MILLISECONDS_OF_MINUTE
	TIME_MILLISECONDS_OF_DAY                                 = 24 * TIME_MILLISECONDS_OF_HOUR
	TIME_SECOND_OF_MINUTE                                    = 60
	TIME_MINUTE_OF_HOUR                                      = 60
	TIME_SECONDS_OF_HOUR                                     = TIME_SECOND_OF_MINUTE * TIME_MINUTE_OF_HOUR
	TIME_SECONDS_OF_DAY                                      = TIME_SECONDS_OF_HOUR * 24
	TIME_YEAR                                                = "yyyy"                    // year
	TIME_SIMPLE_YEAR                                         = "yy"                      // suffix year
	TIME_MONTH                                               = "MM"                      // month
	TIME_DAY                                                 = "dd"                      //day
	TIME_HOUR                                                = "HH"                      //hour
	TIME_MINUTE                                              = "mm"                      //minute
	TIME_SECOND                                              = "ss"                      //second
	TIME_MILLISECOND                                         = "SSS"                     //millisecond
	TIME_YEAR_MONTH                                          = "yyyy-MM"                 // year month
	TIME_YEAR_MONTH_DAY                                      = "yyyy-MM-dd"              // year-month-day
	TIME_YEAR_MONTH_DAY_CN                                   = "yyyy年M月d日"               // year-month-day
	TIME_YEAR_MONTH_DAY_HOUR_MINUTE_SECOND                   = "yyyy-MM-dd HH:mm:ss"     //always used for database
	TIME_YEAR_MONTH_DAY_HOUR_MINUTE_SECOND_MILLISECOND       = "yyyy-MM-dd HH:mm:ss,SSS" //always used for log
	TIME_HOUR_MINUTE_SECOND                                  = "HH:mm:ss"                //hour:minute:second
	TIME_DEFAULT_DATE_FORMAT                                 = "EEE MMM dd HH:mm:ss zzz yyyy"
	TIME_UNION_YEAR_MONTH_DAY_HOUR_MINUTE_SECOND_MILLISECOND = "yyyyMMddHHmmssSSS"
	TIME_UNION_YEAR_MONTH_DAY_HOUR_MINUTE_SECOND             = "yyyyMMddHHmmss"
	TIME_UNION_YEAR_MONTH_DAY_HOUR_MINUTE                    = "yyyyMMddHHmm"
	TIME_UNION_YEAR_MONTH_DAY_HOUR                           = "yyyyMMddHH"
	TIME_UNION_YEAR_MONTH_DAY                                = "yyyyMMdd"
	TIME_UNION_YEAR_MONTH                                    = "yyyyMM"
)
