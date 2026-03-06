package types

type IMEI string

func (i *IMEI) IsValid() bool {
	return true
}
