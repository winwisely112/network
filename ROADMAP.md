## ROADMAP


Really really sorry this is so long. Explaining all this stuff is not fun either...


Like all security its a tradeoff. And a matter of layering to defend
against different aspects.
Also for a project that has really tiny resources we have to find a
happy medium also otherwise we wont ship anything.


Short answer:

- Client P2P wont protect anyone from deep packet inspection.
- Take the client and run it in a emulator and wireshark it and your hosed.
- QUIC is the only thing that wont leak without a VPN because its TCP inside a UDP tunnel that does not leak anything. BUT the IETF is still fighting over if / if not it will leak because the Carriers want to do "traffic management", etc
- VPN in place now gets use full DPI resistance and we can use Server Pub SUB to get going fast !
 - I would rather have my attackers packets inside my VPN where i can watch them and then mitigate actually !
 - Users DONT leak IP or anything to each other because we are starting with Server based PUB SUB.

Long answer:


Below its a bit of a story but its the most honest way to discuss it without just resorting to quoting techncial mumbo jumbo....




I know P2P is where all the hotness is, but I want to highlight is Practicalities:
- Climate emergency is now, so we need to get something in place as soon as possible.
- High quality devs are few and far between and P2P is a little known area and much harder.
- Getting a pure P2P system is over kill for v1 for many other reasons which i try to explain below.
- We can still protect the Users from each other.


We wont do the enrollment "matchmaking" over P2P. It will be classic
RPC actually, and even NOT PUB SUB.
- There is no need and in fact its not the right solution because we
need to do a fair bit of Server logic to do correct matching. This is
because of the business logic.
- I could easily move that logic to the Client and keep the data that
drives the logic up to date on all clients via a PUB SUB pattern but
its just not worth the complexity for V1. For V2 its not hard because we are moving the same code golang from the Server to the Client. So its DRY :) SO what i am saying is here is a great roadmap to do it later without boiling the development ocean.
- So we do not have the problem of protecting Users from each other at
this point !
- At this point we have no idea who they are and do not care. All they
are doing is Campaign shopping as i call it.

BUT they are leaking that they are shopping !. So we really need to sign
them up over traditional HTTPS and
- Get their pub key onto our server.
- Get them onto the VPN so they dont leak.
- The Demo does not have this in there yet !!
 - The Desktop and Mobile versions have this integrated. For Web, they can use installed VPN client that we provide. 

To get on the VPN they just sign up ( in the app) over HTTPS and the pub / priv key dance happens.
- We use their email to do the signup song and dance process. The band is Email - yes i know it sucks but its least worst option...
- So the Gov knows they are using the system. Out of bounds signup is also possible but not practical. SMS is maybe less worse, but like Wire we should not expect them to give us their Mobile number as the Optics are bad, and SMS is totally back doored at Network Carrier level anyway. Telecom ENUM system is a joke. SMS spoofing is too easy. 
- Interesting whats happening with Google Verified SMS though - seems to work by sending SMS's under the hood client to client and the code on the device watching for it arriving and so building up a Verified Graph that is sent back to Google HQ. Simple stuff really. I hope they make this transparent. But anyway not useful for us right now and its Android only.  But i am thinking about borrowing this approach actually internally for User Verification against Bad Actors building wireshark harnesses / bots against us to automate the VPN endpoint discovery and other attack vectors.

Their private key stays on their device and when they add
a new device they transfer it via QR codes out of band.

Once they are on our VPN we are protecting them leaking to
their ISP and DNS.

Now we want to protect them leaking to each other...
- Server PUB SUB does that and thats why for V1 its best. 

Now we take them through the enrollment and they campaign shop.
- Its all classic HTTP RPC here.
- All data and meta data about them is encrypted on our servers
using their own pub key. Its useless to an attacker.
- We only have to hold the relayed PUB SUB data. Their Client holds it all for ever.
  - The QR code transfer can also bootstrap the data from one of their devices to another. Have this code working already doing pure P2P transfer of data without needing us. So its a different architectural approach fundamentally.
- A Bad actor that has our encrypted server data has very little. A post MITM attack vector will be pointless.

Now they start to Collaborate using the tools to work up a campaign together.
- These are all the collaboration modules like Docs, Maps and geo fencing, Cal, and
Video conferencing.
- This is all 100% PUB SUB using classic CDC ( Change Data Capture)
Server architecture. This allows use head room to make the usability
really high.
- This stuff works NOW client and server but we just need to get it integrated.

PUB SUB via Server.
Now whenever data needs to be exchanged between users we use PUB SUB.
On our Servers all User related data is encrypted against the keys the
user generated at Signup point.
So if law enforcement asks for data they will get encrypted data. We
dont have the keys to decrypt it. Only the client does.
- This is what Telegram does too.



DDOS
- The VPN will be attacked. We plan to cycle the endpoints / addresses
of the VPN servers. The VPN is designed exactly for this as a global mesh where 100's of VPN Servers can be added / subtracted constantly.
- Because its PUB SUB the traffic is not high. Clients cache and its encrypted at REST.
- VPN endpoint discovery will be attacked and blocked. So we also plan
to use a back channel like Telegram, Tor. Still working on this.
  - Access to this backchannel is INSIDE the app as well as the VPN
Client. Yes its security via obfuscation in this case.
- The tsunami northern spain app used a similar approach to hide their
proxy and they did not get taken offline. I reverse engineered their
code and the code is useless. They are using Flutter and Golang just like us client side.


Internal DDOS
- They cant get into the VPN unless they are signed in and we have
their public key, so we run metrics to detect doing naughty things and kick them off.
- Other things can be used like rate limiting until they are trusted over a longer period.  
- Already have this in place but not automated.

How the burning question:
"plan to deal with aggregation and localized
matchmaking in an end-to-end encrypted peer-to-peer scenario"
- Its YAGNI IMHO, in that we can achieve the same result with much higher perf and UX experience and less Dev risk using VPN Protected PUB SUB.
- IPFS / LibP2P and Protocol labs are doing all that. I dont
want to use that approach yet as its "boiling the ocean" IMHO. We can
easily bring it in later though because we can run golang embedded in
the clients ( web, Desktop and Mobile). But I want to get boring PUB
SUB via Server going for stage 1 to get a product out and then get
fancy with pure P2P IPFS fun and games later... THE VPN for stage 1 is more important.
- We will use Signal protocol and also PAKE to do the 2PC messaging later.
- PAKE diffie hellman for security session as well as PUB SUB Public /
Private key checking gives 2 phases to this. All this is already available as code. I just dont want to bring it in until phase 2 or 3 !!
- We already use PAKE for synchronising data from one users device to another !! 
 - This is how a user enrols another of their devices.
 - Pub SUB then makes sure all the users devices are kept up to date. No fancy developer skills needed and super low risk.  
 - They can blow away a device at any time and all is fine.
- This stuff works now and we have code for it in golang from the many libs out there. 
- BUT better to use Server PUB SUB for stage 1 to get a Product out the door !!

Things in place to allow actual phasing up whilst in Production:

- We are doing feature flagging and so can introduce these riskier archi approaches on 1% of user base to try them out. OPS Telemetry will tell us how that goes. And users can as we have decent Crash Detection in place.
- The feature flag server is not up yet. Soon once i get more resource.
- Build is local and CI / CD using Docker and K8 / Helm based so we can scale...
- The True P2P code works but we need tooling to make it workable in practice at DEV and OPS level. This is being worked on ..
- The Flutter client will be extended with embedded golang to make the pure P2P work.
- It is all 100% golang and so easy to test. Try testing client code thats is designed to only run in an embedded environment :) The golang client can run in a Terminal outside a Mobile and so easy peasy to test. You need to test for P2P PUB SUB because of the race conditions in this area where state is eventually consistent and you have network partitions happening all the time (reality of P2P).
- You cant do P2P Video conferencing without WebRTC or QUIC ANYWAY and need STUN and TURN Servers anyway because only 90% of users worldwide can use WebRTC or QUIC because the other 10% are blocked by Symmetric Routers anyway !!
- Video Conf is important for enrollment verification we have found from Due Diligence with XR and others.
- You can do Video Conferencing simulcast from the Servers though with SDP, which also gives you multi party. This also allows us to do Live Broadcasting via a backchannel into Youtube which is useful from a Feature POV. So you need Servers anyway is my point, and protecting them and users via VPN makes sense then.


- So get the Server PUB SUB approach done with Webrtc video conf in place and do the fancy P2P Pub SUB stuff in a later phase using feature flagging is my point.


Lastly i want to add about me and Optics.
- I often work for governments ( ofen related to security ) as a consultant and have various contacts from that. 
- For example the German gov has 47 developers just working on breaking
Telegram, and still not succeeding ( so i am told ) because they are attacking an N hard problem and wont ship it ever....
- This is how i know about the German Gov and their project to crack Telegram.
- So i need to not be on the web pages !!! My access to knowing what is happening will be shot. You cant have your cake and eat it too.

