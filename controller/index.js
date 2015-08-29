var socket = io.connect();

// var map = [];
// onkeydown = onkeyup = function (e) {
// 	e = e || event;
// 	map[e.keyCode] = e.type == "keydown";
// 	if (map[87]) socket.emit("move", "up"); 
// 	if (map[83]) socket.emit("move", "down"); 
// 	if (map[65]) socket.emit("move", "left"); 
// 	if (map[68]) socket.emit("move", "right"); 
// };

document.getElementById('up').onclick = function() {
	socket.emit("move", "up");
}
document.getElementById('down').onclick = function() {
	socket.emit("move", "down");
}
document.getElementById('left').onclick = function() {
	socket.emit("move", "left");
}
document.getElementById('right').onclick = function() {
	socket.emit("move", "right");
}