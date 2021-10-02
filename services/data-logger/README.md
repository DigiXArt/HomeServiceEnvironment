# DATA-LOGGER
data-logger is a service that can be used to store custom json data items
such as logs or sensor data or what ever you want.
To structure data by type this service implments collections that are dynamically
created as soon as data is saved to a collection, specified by name.
It is important to know that data can only queried by collection and timeframe.
Also data can't be deleted, so consider this as a long term storage for immutable
data.
You can use this if you want a very lightweight json data storage for your services.
It shows how you can split data into seperate data files, read query para