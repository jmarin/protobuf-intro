package main

import (
	//"github.com/golang/protobuf/proto"
	"github.com/jmarin/protobuf-intro"
	"log"
)

func main() {
	phone1 := &protobuf_intro.PhoneNumber{
		Number:    "555-5555",
		PhoneType: protobuf_intro.PhoneType_HOME,
	}

	p1 := &protobuf_intro.Person{
		Name:   "John Doe",
		Id:     1,
		Email:  "jdoe@email.com",
		Phones: []*protobuf_intro.PhoneNumber{phone1},
	}

	phone2 := &protobuf_intro.PhoneNumber{
		Number:    "666-6666",
		PhoneType: protobuf_intro.PhoneType_MOBILE,
	}

	p2 := &protobuf_intro.Person{
		Name:   "Jane Smith",
		Id:     2,
		Email:  "jsmith@email.com",
		Phones: []*protobuf_intro.PhoneNumber{phone2},
	}

	addressBook := &protobuf_intro.AddressBook{
		People: []*protobuf_intro.Person{p1, p2},
	}

	log.Print(addressBook)
}
