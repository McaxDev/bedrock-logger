import { DisplaySlotId, ObjectiveSortOrder, system, world } from "@minecraft/server";
import {http, HttpHeader, HttpRequest} from "@minecraft/server-net";

const backendAddr = 'http://logger:8080/set'
const loggerPassword = 'Axolotland'

const ignoredEvents = {
	blocks: [
		'minecraft:air',
		'minecraft:dirt',
		'minecraft:grass_block',
	],
	entities: [],
}

// 在玩家生成时执行的操作
world.afterEvents.playerSpawn.subscribe(event => {

	const player = event.player
	if (event.initialSpawn) { // 进入游戏执行的操作
		
		// 玩家第一次进入游戏时，给予夜视效果和消息提示
		if (!player.hasTag('joined')) {
			player.addEffect('minecraft:night_vision', 20 * 60 * 10, {
				showParticles: false
			})
			world.sendMessage(`欢迎${player.name}首次加入服务器`)
			player.sendMessage(`QQ群号825805175`)
			player.addTag('joined')

		}

	} else { // 复活时执行的操作
		player.resetLevel()
	}
})

// 在服务器启动时的操作
world.afterEvents.worldInitialize.subscribe(_ => {

	// 检查并创建在线时长计分板
	if (!world.scoreboard.getObjective("play_time")) {
		world.scoreboard.setObjectiveAtDisplaySlot(DisplaySlotId.List, {
			objective: world.scoreboard.addObjective("play_time", "§b游戏时长/分钟"),
			sortOrder: ObjectiveSortOrder.Descending
		})
	}
})

// 每一分钟执行的操作
system.runInterval(() => {

	world.getDimension('overworld').runCommandAsync(
		'scoreboard players add @a play_time 1'
	)

	const players = world.getAllPlayers()
	if (players.length !== 0) {
		const locations = []
		players.forEach(player => {
			locations.push({
				player: player.name,
				dimension: player.dimension.id,
				location: player.location,
			})
		})
		SendLogRequest({
			type: "online",
			data: {
				data: locations,
			},
		})
	}
}, 20 * 60)

function SendLogRequest(body) {
	if (backendAddr) {
		const request = new HttpRequest(backendAddr)
		request.timeout = 10
		request.method = 'Post'
		request.headers.push(new HttpHeader(
			'Authorization',
			`Bearer ${loggerPassword}`
		))
		request.body = JSON.stringify(body)
		http.request(request)
	}
}

function BlockCallback(type) {
	return function(event) {
		if (ignoredEvents.blocks.includes(event.block.typeId)) {
			return
		}
		SendLogRequest({
			type: type,
			data: {
				blockId: event.block.typeId,
				player: event.player.name,
				dimension: event.block.dimension.id,
				x: event.block.x,
				y: event.block.y,
				z: event.block.z,
			}
		})
	}
}

world.afterEvents.playerBreakBlock.subscribe(BlockCallback('break'))
world.afterEvents.playerPlaceBlock.subscribe(BlockCallback('place'))
world.afterEvents.playerInteractWithBlock.subscribe(BlockCallback('interact'))

world.afterEvents.entityDie.subscribe(event => {
	SendLogRequest({
		type: 'die',
		data: {
			deadId: event.deadEntity.typeId,
			deadName: event.deadEntity.nameTag,
			killerId: event.damageSource.damagingEntity?.typeId,
			killerName: event.damageSource.damagingEntity?.nameTag,
			dimension: event.deadEntity.dimension,
			x: Math.floor(event.deadEntity.location.x),
			y: Math.floor(event.deadEntity.location.y),
			z: Math.floor(event.deadEntity.location.z),
		}
	})
})

world.afterEvents.chatSend.subscribe(event => {
	SendLogRequest({
		type: 'chat',
		data: {
			player: event.sender.name,
			message: event.message,
	 }
 })
})

world.afterEvents.playerJoin.subscribe(SessionCallback(true))
world.afterEvents.playerLeave.subscribe(SessionCallback(false))

function SessionCallback(isJoin) {
	return function(event) {
		SendLogRequest({
			type: 'session',
			data: {
				player: event.playerName,
				isJoin: isJoin,
			}
		})
	}
}
