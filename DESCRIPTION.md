####**IMPORTANT**
*The design is NOT set in stone, they are just ideas thrown out while
brainstorming and are best viewed as such. They may not be doable, deemed too
difficult or a bad idea, or otherwise discarded. Those which do make it into the
implementation may be vastly different than first suggested, or this project may
be discarded entirely. We are just spit-balling at the moment, and make no
guarentees about any of it.
_<b>We will release a public specification once everything is figured out.</b>_*


##Anonswap
is a simple protocol to implement sharing of arbitrary files in a way
that is resistant to snooping by thirdparty which may be listening to your
traffic (ISP, government, hackers, etc). The idea is pretty simple: clients
connect to one or more public servers (think like dns) that function shomething
like a community bulletin board like the one found at universities, or other
places where items that are for sale or announcements are thumb-tacked up with
the posters contact details, but papers and thumbtacks are replaced by "records"
(described below), and the posters telephone number is replaced by an IP address
(or other resolvable endpoint) and a public key.

###*_Record_*
:

In addition to a resolvable address where the poster accepts connections and
their public key, records contain a list files which the poster has, and is
willing to share with others, which may have constraints associated with their
access. These constraints might be a certain upload bandwith or count, a
reputation score, speed of upload or download, or a plethera of other possible
requirements that must be met in order to download. Optionally instead of hard
requirements for access, the poster may choose to instead prioritize based on
said constraints, letting client who meet them have priority to their access,
but not denying access to those who don't meet those constraints.

###*_Record Format_*
format:

*TODO*

###*_Transfer Protocol_*
:

When a qualified (meets all set contraints) client has decided they would like
to download one or more files which are shared by a host, they send an __Access
Request__ signed with the requesting clients private key to the host. This
request contains the names of the files (possibly along with their hashes), the
requesters public key, the server that the requester found the listing, and any
timeouts, deadlines, or share protocols that they would like to use for this
transfer. The declared server location is checked to make sure that the sharer
did infact list there, the expiration date is correct and not reached.
If everything checks out, the host assembles the content as laid out under
__Share Initialization__, and sets up the transfer using the requested
parameters. Before sending the content the host sends __Status Report__ message
to the client. This message is described in detail below, but it is basically
an ack/noack message either confirming that all is well and the transfer is
about to start, or reporting the transfer is not happening, and the reason for
that, and possibly suggestions to remedy the error. The main reason for doing
this instead of just starting the transfer is so that the traffic will look
similar whether there was an error or not, to hopfully make it more difficult
for would-be attackers to derive any information from the system.

### *_Access Request_*
format:

*TODO*


###*_Share Initialization_*
:

After the status report is sent (may happen in parallel), the requested files
are packaged together into an encrypted archive, deriving a shared key from the
requesters public key, and the hosts private key. A SHA2 hash is calculated,
which then becomes the name of the archive.
It is encouraged to bundle multiple files into the archive, partly to add some
sort of entropy making individual file identitication more difficult.

If multiple providers have the same files listed in the access
request, different chunks may be downloaded from those providers in parallel and
re-assembled on the client machine. The chunk sizes and provider limits should
be configurable.
