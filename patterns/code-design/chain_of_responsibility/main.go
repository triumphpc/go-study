package main

import "fmt"

type department interface {
	execute(*patient)
	setNext(department)
}

type Handler struct {
	Args string
}

type reception struct {
	Handler

	next department
}

func (r *reception) execute(p *patient) {
	if p.registrationDone {
		fmt.Println("Patient registration already done")
		r.next.execute(p)
		return
	}
	fmt.Println("Reception registering patient")
	p.registrationDone = true
	r.next.execute(p)
}

func (r *reception) setNext(next department) {
	r.next = next
}

type doctor struct {
	Handler

	next department
}

func (d *doctor) execute(p *patient) {
	if p.doctorCheckUpDone {
		fmt.Println("Doctor checkup already done")
		d.next.execute(p)
		return
	}
	fmt.Println("Doctor checking patient")
	p.doctorCheckUpDone = true
	d.next.execute(p)
}

func (d *doctor) setNext(next department) {
	d.next = next
}

type medical struct {
	Handler

	next department
}

func (m *medical) execute(p *patient) {
	if p.medicineDone {
		fmt.Println("Medicine already given to patient")
		m.next.execute(p)
		return
	}
	fmt.Println("Medical giving medicine to patient")
	p.medicineDone = true
	m.next.execute(p)
}

func (m *medical) setNext(next department) {
	m.next = next
}

type cashier struct {
	next department
}

func (c *cashier) execute(p *patient) {
	if p.paymentDone {
		fmt.Println("Payment Done")
		return
	}

	fmt.Println("Cashier getting money from patient patient")
	p.paymentDone = true

	c.next.execute(p)

}

func (c *cashier) setNext(next department) {
	c.next = next
}

type patient struct {
	name              string
	registrationDone  bool
	doctorCheckUpDone bool
	medicineDone      bool
	paymentDone       bool
}

func main() {
	doctor := &doctor{}
	cashier := &cashier{}
	cashier.setNext(doctor)

	//Set next for medical department
	medical := &medical{}
	medical.setNext(cashier)

	//Set next for doctor department

	doctor.setNext(medical)

	//Set next for reception department
	reception := &reception{}
	reception.setNext(doctor)
	reception.Args = "TEST"

	patient := &patient{name: "abc"}

	//Patient visiting
	reception.execute(patient)
}
