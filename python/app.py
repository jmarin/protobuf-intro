import addressbook_pb2

addressbook = addressbook_pb2.AddressBook()
p1 = addressbook.people.add()
p1.id = 1
p1.name = "John Doe"
p1.email = "jdoe@email.com"
phone1 = p1.phones.add()
phone1.number = "555-5555"
phone1.phoneType = addressbook_pb2.HOME

p2 = addressbook.people.add()
p2.id = 2
p2.name = "Jane Smith"
p2.email = "jsmith@email.com"
phone2 = p2.phones.add()
phone2.number = "666-6666"
phone2.phoneType = addressbook_pb2.MOBILE

print addressbook