<template>
	<view>
		<input type="text" v-model="content" />
		<button @click="send()">发送</button>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				user_id: 0,
				dist_id: 0,
				content: "",
				msgType:{
					user_id: 0,
					dist_id: 0,
					content: "",
					msg_type: 0,
					cmd: 0
				}
			}
		},
		methods: {
			createConn(){
				// 创建链接
				uni.connectSocket({
					url: 'ws://127.0.0.1:8088/api/v1/chan?user_id='+this.user_id,
					header: {
						'content-type': 'application/json'
					},
					method:"GET"
				});
				uni.onSocketOpen(function (res) {
				  console.log('WebSocket连接已打开！');
				});
				uni.onSocketError(function (res) {
				  console.log('WebSocket连接打开失败，请检查！');
				});
				
				// 接收消息
				uni.onSocketMessage(function (res) {
				  console.log('收到服务器内容：' + res.data);
				});
			},
			send(){
				this.msgType = {
					user_id: parseInt(this.user_id),
					dist_id: parseInt(this.dist_id),
					content: this.content,
					msg_type: 1,
					cmd: 1,
				}
				console.log(this.msgType);
				// 发送消息
				uni.sendSocketMessage({
				      data: JSON.stringify(this.msgType)
				});
			}
		},
		onLoad(options) {
			this.user_id = options.user_id
			this.dist_id = options.dist_id
		},
		mounted() {
			this.createConn()
		}
	}
</script>

<style>

</style>
