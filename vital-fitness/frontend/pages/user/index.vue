<template>
	<view class="container">
		<!-- 顶部导航栏 -->
		<u-navbar title="个人中心" fixed placeholder />

		<!-- 用户信息卡片 -->
		<view class="user-card">
			<view class="user-avatar">
				<u-avatar
					:src="userInfo.avatar || '/static/images/default-avatar.png'"
					size="large"
					@click="changeAvatar"
				></u-avatar>
				<u-icon name="camera" size="20" color="#fff" class="camera-icon"></u-icon>
			</view>

			<view class="user-info">
				<view class="username-row">
					<text class="username">{{ userInfo.nickname || userInfo.username }}</text>
					<u-icon name="edit-pen" size="18" color="#999" @click="editProfile"></u-icon>
				</view>

				<view class="user-stats">
					<view class="stat-item">
						<text class="stat-value">{{ userInfo.days_tracked }}</text>
						<text class="stat-label">记录天数</text>
					</view>
					<view class="stat-item">
						<text class="stat-value">{{ userInfo.total_workouts }}</text>
						<text class="stat-label">训练次数</text>
					</view>
					<view class="stat-item">
						<text class="stat-value">{{ userInfo.achievements }}</text>
						<text class="stat-label">成就</text>
					</view>
				</view>

				<view class="user-bio" v-if="userInfo.bio">
					<text>{{ userInfo.bio }}</text>
				</view>
			</view>
		</view>

		<!-- 功能菜单 -->
		<view class="menu-section">
			<u-cell-group>
				<u-cell
					title="个人资料"
					icon="account"
					isLink
					@click="toProfile"
				></u-cell>
				<u-cell
					title="健康目标"
					icon="target"
					isLink
					@click="toGoals"
				></u-cell>
				<u-cell
					title="数据统计"
					icon="pie-chart"
					isLink
					@click="toStatistics"
				></u-cell>
				<u-cell
					title="成就系统"
					icon="medal"
					isLink
					@click="toAchievements"
				></u-cell>
			</u-cell-group>
		</view>

		<!-- 设置菜单 -->
		<view class="menu-section">
			<u-cell-group>
				<u-cell
					title="通知设置"
					icon="bell"
					isLink
					@click="toNotifications"
				></u-cell>
				<u-cell
					title="隐私设置"
					icon="lock"
					isLink
					@click="toPrivacy"
				></u-cell>
				<u-cell
					title="通用设置"
					icon="setting"
					isLink
					@click="toSettings"
				></u-cell>
			</u-cell-group>
		</view>

		<!-- 帮助与反馈 -->
		<view class="menu-section">
			<u-cell-group>
				<u-cell
					title="帮助中心"
					icon="help"
					isLink
					@click="toHelp"
				></u-cell>
				<u-cell
					title="意见反馈"
					icon="edit-pen"
					isLink
					@click="toFeedback"
				></u-cell>
				<u-cell
					title="关于我们"
					icon="info-circle"
					isLink
					@click="toAbout"
				></u-cell>
			</u-cell-group>
		</view>

		<!-- 退出登录 -->
		<view class="logout-section">
			<u-button type="error" plain @click="logout">退出登录</u-button>
		</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				userInfo: {
					username: 'fitness_user',
					nickname: '健身达人',
					avatar: '',
					bio: '坚持健身，享受生活',
					days_tracked: 42,
					total_workouts: 128,
					achievements: 15
				}
			}
		},
		methods: {
			changeAvatar() {
				// 更换头像
				uni.chooseImage({
					count: 1,
					success: (res) => {
						// 这里应该上传头像到服务器
						this.userInfo.avatar = res.tempFilePaths[0]
						uni.showToast({
							title: '头像更新成功',
							icon: 'success'
						})
					}
				})
			},

			editProfile() {
				uni.navigateTo({
					url: '/pages/user/profile'
				})
			},

			toProfile() {
				uni.navigateTo({
					url: '/pages/user/profile'
				})
			},

			toGoals() {
				uni.navigateTo({
					url: '/pages/user/goals'
				})
			},

			toStatistics() {
				uni.navigateTo({
					url: '/pages/statistics/index'
				})
			},

			toAchievements() {
				uni.navigateTo({
					url: '/pages/user/achievements'
				})
			},

			toNotifications() {
				uni.navigateTo({
					url: '/pages/user/notifications'
				})
			},

			toPrivacy() {
				uni.navigateTo({
					url: '/pages/user/privacy'
				})
			},

			toSettings() {
				uni.navigateTo({
					url: '/pages/user/settings'
				})
			},

			toHelp() {
				// 打开帮助中心
				uni.showToast({
					title: '帮助中心',
					icon: 'none'
				})
			},

			toFeedback() {
				uni.navigateTo({
					url: '/pages/user/feedback'
				})
			},

			toAbout() {
				uni.navigateTo({
					url: '/pages/user/about'
				})
			},

			logout() {
				uni.showModal({
					title: '确认退出',
					content: '确定要退出登录吗？',
					success: (res) => {
						if (res.confirm) {
							this.$store.dispatch('logout')
							uni.redirectTo({
								url: '/pages/auth/login'
							})
						}
					}
				})
			}
		}
	}
</script>

<style lang="scss" scoped>
.container {
	padding: 20rpx;
}

.user-card {
	background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
	border-radius: 20rpx;
	padding: 40rpx 30rpx;
	margin-bottom: 30rpx;
	color: white;
	position: relative;
	overflow: hidden;

	&::before {
		content: '';
		position: absolute;
		top: -50%;
		left: -50%;
		width: 200%;
		height: 200%;
		background: radial-gradient(circle, rgba(255,255,255,0.1) 0%, transparent 70%);
		transform: rotate(30deg);
	}

	.user-avatar {
		display: flex;
		justify-content: center;
		margin-bottom: 30rpx;
		position: relative;

		.camera-icon {
			position: absolute;
			bottom: 0;
			right: 20%;
			background: #3c9cff;
			border-radius: 50%;
			padding: 5rpx;
		}
	}

	.user-info {
		text-align: center;

		.username-row {
			display: flex;
			align-items: center;
			justify-content: center;
			margin-bottom: 20rpx;

			.username {
				font-size: 36rpx;
				font-weight: bold;
				margin-right: 10rpx;
			}
		}

		.user-stats {
			display: flex;
			justify-content: space-around;
			margin-bottom: 30rpx;

			.stat-item {
				text-align: center;

				.stat-value {
					display: block;
					font-size: 32rpx;
					font-weight: bold;
					margin-bottom: 5rpx;
				}

				.stat-label {
					font-size: 24rpx;
					opacity: 0.9;
				}
			}
		}

		.user-bio {
			font-size: 26rpx;
			opacity: 0.9;
			line-height: 1.5;
		}
	}
}

.menu-section {
	background: white;
	border-radius: 20rpx;
	margin-bottom: 20rpx;
	overflow: hidden;
}

.logout-section {
	padding: 30rpx 0;

	/u-button {
		border-radius: 10rpx;
	}
}
</style>