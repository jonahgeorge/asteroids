function Ship() {}

Ship.draw = function(ship) {

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
