
# Seat Holding
for every booking we create a key in redis with:

key: booking:movieid:seatid
val: booking data

this has a TTL to ensure no one else can book for that particular time ensuring no double booking occurs


We also create another key for reverse lookup of the booking. Each booking has an unique Id , we map each booking to the session Id ie booking Id

key: sessionId (session:bookingId)
val: key (booking:movieId:sessionId)

This has the same TTL to ensure that it expires if booking is not confirmed;

# Seat Confirmation