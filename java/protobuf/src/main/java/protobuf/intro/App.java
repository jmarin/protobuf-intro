package protobuf.intro;

import java.util.ArrayList;
import java.io.FileOutputStream;
import java.io.FileInputStream;
import java.io.FileNotFoundException;
import java.io.IOException;

public class App {
    public static void main( String[] args ) {
        System.out.println( "Protobuf example" );

        Addressbook.PhoneNumber phone1 = Addressbook.PhoneNumber.newBuilder()
                                .setNumber("555-5555")
                                .setPhoneType(Addressbook.PhoneType.HOME)
                                .build();

        Addressbook.Person john = Addressbook.Person.newBuilder()
                .setId(1)
                .setName("John Doe")
                .setEmail("jdoe@email.com")
                .addPhones(phone1)
                .build();

        Addressbook.PhoneNumber phone2 = Addressbook.PhoneNumber.newBuilder()
                .setNumber("666-6666")
                .setPhoneType(Addressbook.PhoneType.MOBILE)
                .build();

        Addressbook.Person jane = Addressbook.Person.newBuilder()
                .setId(2)
                .setName("Jane Smith")
                .setEmail("jsmith@email.com")
                .addPhones(phone2)
                .build();

        Addressbook.AddressBook addressBook = Addressbook.AddressBook.newBuilder()
                .addPeople(john)
                .addPeople(jane)
                .build();

        try {
            FileOutputStream output = new FileOutputStream("target/addressbook.pbf");
            addressBook.writeTo(output);
            FileInputStream input = new FileInputStream("target/addressbook.pbf");
            Addressbook.AddressBook.Builder addressBook2 = Addressbook.AddressBook.newBuilder();
            addressBook2.mergeFrom(input);
            System.out.println(addressBook2);
        } catch (FileNotFoundException e){
            System.out.println(e.getLocalizedMessage());
        } catch (IOException ioe) {
            System.out.println(ioe.getLocalizedMessage());
        }

    }
}
