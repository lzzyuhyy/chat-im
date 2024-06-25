<template>
	<view>
		<button  @click="toFriendList">登录</button>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				
			}
		},
		methods: {
			toFriendList(){
				
				
				
				
				uni.navigateTo({
					url:"/pages/friendList/friendList"
				})
			}
		}
	}
</script>

<style>

</style>
