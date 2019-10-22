package service

import (
	"strconv"
)

type Date struct{
	Year, Month, Day, Hour, Minute int
}

func isValid(date Date) bool{
	if date.Year < 1000 || date.Year > 9999 || date.Month > 12 || date.Month < 1 || date.Hour > 24 || date.Hour < 0 || date.Minute > 60 || date.Minute < 0 || date.Day < 1{
		return false
	}
	switch date.Day {
	case 1, 3, 5, 7, 8, 10, 12:
		if date.Day > 31 {
			return false
		}
	case 2:
		if date.Year % 4 == 0{
			if date.Day > 29{
				return false
			}
		}else{
			if date.Day > 28{
				return false
			}
		}
	default:
		if date.Day > 30{
			return false
		}
	}
	return true
}

func StringToDate(dString string) Date{
	res := Date{
		Year: 	0,
		Month:  0,
		Day:    0,
		Hour:   0,
		Minute: 0,
	}
	if len(dString) != 16 || dString[4] != '-' || dString[7] != '-' || dString[10] != '-' || dString[13] != ':'{ //date format is xxxx-xx-xx-xx:xx
		return res
	}
	tmp := res
	res.Year, _ = strconv.Atoi(dString[0:4])
	res.Month, _ = strconv.Atoi(dString[5:7])
	res.Day, _ = strconv.Atoi(dString[8:10])
	res.Hour, _ = strconv.Atoi(dString[11:13])
	res.Minute, _ = strconv.Atoi(dString[14: 16])
	if isValid(res){
		return res
	}else{
		return tmp
	}
}

func DateToString(date Date) string{
	res := strconv.Itoa(date.Year) + "-" + strconv.Itoa(date.Month) + "-" + strconv.Itoa(date.Day) + "-" + strconv.Itoa(date.Hour) + "-" + strconv.Itoa(date.Minute)
	return res
}

func compare_date(date1, date2 Date) bool{ // return date1 > date2 ? true : false
	if date1.Year > date2.Year{
		return true
	}else if date1.Year < date2.Year{
		return false
	}else{
		if date1.Month > date2.Month {
			return true
		}else if date1.Month < date2.Month{
			return false
		}else{
			if date1.Day > date2.Day{
				return true
			}else if date1.Day < date2.Day{
				return false
			}else{
				if date1.Hour > date2.Hour{
					return true
				}else if date1.Hour < date2.Hour {
					return false
				}else{
					return date1.Minute > date2.Minute
				}
			}
		}
	}
}

