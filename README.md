# email-address-validator
Validating email addresses should be free and open-source

## how it's done
Asks the user for the domain name of the email (gmail.com). Then parses through the net pacakge on Go in search of MX (mail exchange), SPF (sender policy framework), DMARC (Domain-based Message Authentication Reporting and Conformance) records. 

## for my first time learners
For context. MX records are like address books for computers. If it exists, then there's a registered address, meaning it's most likely legitimate. Moving on, SPF records are like special passwords that tell computers which email senders are allowed to send messages on behalf of a domain. Finally, DMARC records are like security guards that help protect emails and make sure they come from the right place. Kinda like a secret code that only trusted messengers can use.
