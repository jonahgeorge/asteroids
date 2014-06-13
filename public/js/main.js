var canvas = document.getElementById("game");
	canvas.height = window.innerHeight;
	canvas.width = window.innerWidth;

var context = canvas.getContext("2d");

var socket = io.connect();
var players = [];

var map = [];
onkeydown = onkeyup = function (e) {
	e = e || event;
	map[e.keyCode] = e.type == "keydown";
	if (map[87]) socket.emit("move", "up"); 
	if (map[83]) socket.emit("move", "down"); 
	if (map[65]) socket.emit("move", "left"); 
	if (map[68]) socket.emit("move", "right"); 
};

socket.on('update', function(data){
	players = JSON.parse(data);
});

setInterval(function() {
	context.clearRect(0, 0, canvas.width, canvas.height);	
	players.forEach(function (player) {
		Ship.draw(player.ship);
	});
}, 1000 / 30);

