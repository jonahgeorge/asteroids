var express = require('express')
var app = express();
var http = require('http').Server(app);
var io = require('socket.io')(http);

var Ship = require('./client.js').Ship;
var Point = require('./client.js').Point;

var clients = {};

app.use('/controller', express.static('../controller'));
app.use('/observer', express.static('../observer'));

setInterval(function() {
    for (var key in clients) {
		// Update location
		clients[key].position.x += clients[key].velocity.x
		clients[key].position.y += clients[key].velocity.y
    }

	io.sockets.emit('update', JSON.stringify(clients));
}, 1000 / 30);

io.on('connection', function(socket){
  console.log('a user connected');

	clients[socket.id] = new Ship(
		new Point(200, 200),
		new Point(0, 0),
		new Point(0, 0),
		0,
		1,
		0.3,
		getRandomColor()
	);

	socket.on('move', function(command) {
		handleMove(command, socket);
	});
	socket.on('disconnect', function() {
		handleDisconnect(socket);
	});
});

http.listen(3000, function(){
  console.log('listening on *:3000');
});

function handleMove(command, socket) {
	idx = socket.id;

	switch(command) {
	case "up":
		clients[idx].acceleration.x = clients[idx].thrust * Math.sin(clients[idx].heading);
		clients[idx].acceleration.y = clients[idx].thrust * Math.cos(clients[idx].heading);
		clients[idx].velocity.x += clients[idx].acceleration.x;
		clients[idx].velocity.y += clients[idx].acceleration.y;

	case "down":
		// nada

	case "left":
		clients[idx].heading += clients[idx].turnSpeed;

	case "right":
		clients[idx].heading -= clients[idx].turnSpeed;
	}
}

function handleDisconnect(socket) {
	console.log("user " + socket + " disconnected.")
	delete clients[socket];
}

function getRandomColor() {
    var letters = '0123456789ABCDEF'.split('');
    var color = '#';
    for (var i = 0; i < 6; i++ ) {
        color += letters[Math.floor(Math.random() * 16)];
    }
    return color;
}