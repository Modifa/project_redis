package models

/*Returnblock ...  */
type Returnblock struct {
	Status  bool
	Message string
	Data    interface{}
}

func (r *Returnblock) New(status bool, message string, data interface{}) *Returnblock {

	return &Returnblock{
		status,
		message,
		data,
	}

}
