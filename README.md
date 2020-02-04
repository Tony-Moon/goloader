# GoLoader

Go loader is an app to meant to hone my go skills until I can actuallu start working on a project at work.
For now, I am just trying not to loose my touch and get rusty.
I figured downloading/torrenting files should be a fun and interesting challenge.
Fingers crossed!


## How I envision this working

I envision this working on three different applications, for an extra bit of spice thrown into the challenge.
These three apps I will call the sender, the portal and the reciever. 
Pretty self-explanitory right?


### App 1: The Sender

The sender application will be able to upload any file type to the cloud (the portal application). 
The sender will expect two notifications from the portal on the status of the file.
For the first notification:
- If the portal does not send a notification back to the sender, the sender should assume the portal did not successfully recieve the file. 
The sender should not expect a second repsonse from the portal.
- If the portal sends an error to the sender (whether a file corruption or incomple file), the sender should be able to act accordingly.
The sender should not expect a second repsonse from the portal.
- If the portal sends an all clear, the sender should expect another notification.

For the second notification:
- If the portal could not send the file to the reciever it should notify the sender.
- If the portal successfully passes the file to the reciever, it should notify the sender.


### App 2: The Portal

The portal will recieve the file from the sender. 
The portal will then store the file in a sort of temporary storage.
The portal will send up to two notification back to the sender. 
The first notification to the sender will only be affected by the communication between the sender and the portal:
- If the portal recieves a partial or corrupted file, it should delete the file and ask the sender to resend the file. 
The portal will not send a second notification to the sender.
- If the portal fails to store the file (out of memory perhaps) the portals should send an error and notify the sender to not send anything else.
The portal will not send a second notification to the sender. 
- If the portal successfully recieves and stores the file, it should notify the sender of that.
From here, the portal will send a second notification to the sender.

The second notification will tell the sender of the status between the portal and the reciever. 
Essentially, the portal will expect the same kind of response from the reciever as the sender expected from the portal. 
The portal will be expecting a notification from the reciever:
- If the portal does not recieve a notification it should assume the reciever has not recieved the file and act accordingly.
- If the portal recieves a notification from the reciever that tells of an error, the portal should act accordingly.
The portal will send the same error message back to the sender.
- If the portal recieves an all clear, the portal will send another all clear to the sender.

### App 3: The Reciever

Essentailly, the reciever works like half of the portal.
The reciever will be expecting a file.
If the file is corrupted or is impartial, it should notify the portal to try again after deleting the corrupted/incomplete file.
If the file is downloaded correctly, it should send an all clear back to the portal.