
function InitApp() {
	const app = new PIXI.Application({
		width: backgroundWidth,
		height: backgroundHeight,
		// backgroundAlpha: true,
		// transparent : true,
		// backgroundColor: 0x1099bb
	});

	// Setup background

	const backgroundSprite = PIXI.Sprite.from('/assets/beirut_2-After-Hi-res.jpg');
	//const vehicle = PIXI.Sprite.from('/assets/vehicle.png');

	backgroundSprite.anchor.set(0.5);
	backgroundSprite.x = app.screen.width / 2;
	backgroundSprite.y = app.screen.height / 2;
	backgroundSprite.width  = backgroundWidth;
	backgroundSprite.height = backgroundHeight;
	backgroundSprite.interactive = true;

	backgroundSprite.on('pointerdown', (event) => { 
		// console.log("eve", event.data.global);
		console.log("URL", window.location.pathname);
		if (window.location.pathname == "/agent.html") {
			TargetPos.x = event.data.global.x;
			TargetPos.y = event.data.global.y;
		}
	});

	// Whenever a new agent must be added

	app.NewAgent = function(pos) {
		let agent = PIXI.Sprite.from('/assets/vehicle.png');

		agent.anchor.set(0.5);

		// move the sprite to the center of the screen
		// vehicle.x = app.screen.width / 2;
		// vehicle.y = app.screen.height / 2;
		agent.x = pos.x;
		agent.y = pos.y;

		agent.scale.x = 0.25;
		agent.scale.y = 0.25;

		this.stage.addChild(agent);

		return agent;
	}

	app.stage.addChild(backgroundSprite);

	return app;
}
