package main

import "fmt"

/* https://golangbyexample.com/chain-of-responsibility-design-pattern-in-golang/

When to Use:
The pattern is applicable when there are multiple candidates to process the same request.
When you don’t want the client to choose the receiver as multiple objects can
handle the request. Also, you want to decouple the client from receivers. The
Client only needs to know the first element in the chain.
    -As in the example of the hospital, a patient first goes to the reception and then
	reception based upon a patient’s current status sends up to the next handler in the
	chain.
*/

type Patient struct {
	name              string
	registrationDone  bool
	doctorCheckUpDone bool
	medicineDone      bool
	paymentDone       bool
}

// Handler interface
type Department interface {
	execute(*Patient)
	setNext(Department)
}

// Reception
// concrete handler
type Reception struct {
	next Department
}

func (r *Reception) execute(p *Patient) {
	if p.registrationDone {
		fmt.Println("Patient registration already done")
		r.next.execute(p)
		return
	}
	fmt.Println("Reception registering patient")
	p.registrationDone = true
	r.next.execute(p)
}

func (r *Reception) setNext(next Department) {
	r.next = next
}

// Doctor
type Doctor struct {
	next Department
}

func (d *Doctor) execute(p *Patient) {
	if p.doctorCheckUpDone {
		fmt.Println("Doctor checkup already done")
		d.next.execute(p)
		return
	}
	fmt.Println("Doctor checking patient")
	p.doctorCheckUpDone = true
	d.next.execute(p)
}

func (d *Doctor) setNext(next Department) {
	d.next = next
}

// Medical
type Medical struct {
	next Department
}

func (m *Medical) execute(p *Patient) {
	if p.medicineDone {
		fmt.Println("Medicine already given to patient")
		m.next.execute(p)
		return
	}
	fmt.Println("Medical giving medicine to patient")
	p.medicineDone = true
	m.next.execute(p)
}

func (m *Medical) setNext(next Department) {
	m.next = next
}

// Cashier
type Cashier struct {
	next Department
}

func (c *Cashier) execute(p *Patient) {
	if p.paymentDone {
		fmt.Println("Payment Done")
	}
	fmt.Println("Cashier getting money from patient patient")
}

func (c *Cashier) setNext(next Department) {
	c.next = next
}

// 顺序类似于wire_gen的实例化顺序
func main() {
	cashier := &Cashier{}

	// Set next for medical department
	medical := &Medical{}
	medical.setNext(cashier)

	// Set next for doctor department
	doctor := &Doctor{}
	doctor.setNext(medical)

	// Set next for reception department
	reception := &Reception{}
	reception.setNext(doctor)

	patient := &Patient{name: "abc"}
	// Patient visiting
	reception.execute(patient)
}
