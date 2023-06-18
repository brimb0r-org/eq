db.eq.remove({})
for (let i = 0; i < 1000000; i++) {
    db.eq.insert({
        "_id": "5d8a342ba9e109cb56" + i,
        "Name": "Eq",
        "lastupdated": ISODate("2020-09-23T12:55:23.902Z"),
        "Activity": "Spend",
        "Publish_status": false
    });
}