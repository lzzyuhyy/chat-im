<template>
	<view>
		<view>
			<scroll-view :scroll-top="scrollTop" scroll-y="true" class="scroll-Y" @scrolltoupper="upper"
				@scrolltolower="lower" @scroll="scroll">
				<ul>
					<li v-for="(item, index) in friendList" :key="item.ID" class="contact-item" @click="toChat(item.owner_id, item.dist_id)">
						<img :src="item.user_info.avatar" alt="" style="width: 50px;" class="contact-avatar" />
						<p class="contact-nickname">{{item.user_info.nickname}}</p>
						<p class="contact-last-message-time">{{getDate(item.CreatedAt)}}</p>
					</li>
				</ul>
			</scroll-view>
		</view>
		<view @tap="goTop" class="uni-link uni-center uni-common-mt toTop">
			<uni-icons type="arrow-up" size="30" color="#c3c3c3"></uni-icons>
		</view>


	</view>
</template>

<script>
	export default {
		data() {
			return {
				scrollTop: 0,
				old: {
					scrollTop: 0
				},
				friendList: []
			}
		},
		methods: {
			toChat(userId,distId){
				uni.navigateTo({
					url:"/pages/chat/chat?user_id="+userId+"&dist_id="+distId
				})
			},
			getDate(date) {
				let newDate = new Date(date)
				let hours = String(newDate.getHours()).padStart(2, '0'); // 两位数小时  
				let minutes = String(newDate.getMinutes()).padStart(2, '0'); // 两位数分钟  
				return `${hours}:${minutes}`;
			},
			getFriendList() {
				uni.request({
					url: "api/v1/user/friend/list",
					data: {
						user_id: 1,
					},
					success: (res) => {
						console.log(res.data);
						if (res.data.code == 0) {
							this.friendList = res.data.data.friend_list
						} else {
							alert(res.data.messgae)
						}
					},
					fail: (e) => {
						alert(e)
					}
				})
			},
			upper: function(e) {
				console.log(e)
			},
			lower: function(e) {
				console.log(e)
			},
			scroll: function(e) {
				console.log(e)
				this.old.scrollTop = e.detail.scrollTop
			},
			goTop: function(e) {
				// 解决view层不同步的问题
				this.scrollTop = this.old.scrollTop
				this.$nextTick(function() {
					this.scrollTop = 0
				});
				// 提示信息
				uni.showToast({
					icon: "none",
					title: "纵向滚动 scrollTop 值已被修改为 0"
				})
			}
		},
		mounted() {
			this.getFriendList()
		}
	}
</script>

<style>
	.scroll-Y ul {
		list-style: none;
		padding: 0;
		margin: 0;
	}

	.contact-item {
		position: relative;
		width: 100%;
		height: 70px;
		border-bottom: 1px solid #c3c3c3;
	}

	.contact-avatar {
		width: 50px;
		height: 50px;
		border-radius: 50%;
		position: absolute;
		top: 10px;
		left: 8px;
	}

	.contact-nickname {
		/* font-size: ; */
		position: absolute;
		top: 10px;
		left: 68px;
	}

	.contact-last-message-time {
		font-size: 12px;
		color: #ccc;
		position: absolute;
		top: 10px;
		right: 8px;
	}

	.toTop {
		width: 40px;
		height: 40px;
		text-align: center;
		line-height: 40px;
		background-color: aliceblue;
		border-radius: 50%;
		position: fixed;
		bottom: 5px;
		right: 5px;
	}
</style>