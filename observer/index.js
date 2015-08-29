var canvas = document.getElementById("game");
	canvas.height = window.innerHeight;
	canvas.width = window.innerWidth;

var context = canvas.getContext("2d");

var socket = io.connect();
var players = [];

socket.on('update', function(data){
	players = JSON.parse(data);
});

setInterval(function() {
	context.clearRect(0, 0, canvas.width, canvas.height);	
    for (var key in players) {
    	drawShip(players[key]);
    }
}, 1000 / 30);

function drawShip(ship) {
	var arm1 = 20;
	var arm2 = 10;

	context.beginPath();

	context.moveTo(
		ship.position.x + (arm1*Math.sin(ship.heading)), 
		ship.position.y + (arm1*Math.cos(ship.heading)));

	context.lineTo(
		ship.position.x + (arm2*Math.sin(ship.heading+90)), 
		ship.position.y + (arm2*Math.cos(ship.heading+90)));

	context.lineTo( 
		ship.position.x + (arm2*Math.sin(ship.heading-90)), 
		ship.position.y + (arm2*Math.cos(ship.heading-90)));

	context.fillStyle = ship.color;
	context.fill();
}