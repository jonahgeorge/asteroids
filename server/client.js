exports.Point = function (x, y) {
	this.x = x;
	this.y = y;
}

exports.Ship = function (position, acceleration, velocity, heading, thrust, turnSpeed, color) {
	this.position = position;
	this.acceleration = acceleration;
	this.velocity = velocity;
	this.heading = heading;    
	this.thrust = thrust;     
	this.turnSpeed = turnSpeed;
	this.color = color;
}