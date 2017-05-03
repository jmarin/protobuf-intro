var protobuf = require('protobufjs');

console.log("Protocol Buffers in JavaScript");

protobuf.load("../protobuf/addressbook.proto", function(err, root){
  if (err)
    throw err;

  var PhoneMessage = root.lookupType("protobuf.intro.PhoneNumber");
  var PersonMessage = root.lookupType("protobuf.intro.Person");
  var AddressBookMessage = root.lookupType("protobuf.intro.AddressBook");

  var f1 = { number: "555-5555", phoneType: "HOME"}
  var phone1 = PhoneMessage.create(f1);

  var p1 = { name: "John Doe", id: 1, email: "jdoe@email.com", phones: [phone1]}

  var person1 = PersonMessage.create(p1);

  var f2 = { number: "666-6666", phoneType: "MOBILE"}
  var phone2 = PhoneMessage.create(f2);

  var p2 = { name: "Jane Smith", id: 2, email: "jsmith@email.com", phones: [phone2]}

  var person2 = PersonMessage.create(p2);

  var payload = { people: [person1, person2] }
  var addressbook = AddressBookMessage.create(payload);

  console.log(addressbook);

  console.log("Protocol Buffer in binary format:\n");

  console.log(AddressBookMessage.encode(addressbook).finish());

});