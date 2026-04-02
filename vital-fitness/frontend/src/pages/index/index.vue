<template>
	<view class="container">
		<!-- 顶部导航栏 -->
		<u-navbar title="Vital Fitness" fixed placeholder>
			<view slot="right">
				<u-icon name="setting" size="20" @click="toSetting"></u-icon>
			</view>
		</u-navbar>

		<!-- 用户信息卡片 -->
		<view class="user-card">
			<view class="user-info">
				<u-avatar :src="userInfo.avatar || '/static/images/default-avatar.png'" size="large"></u-avatar>
				<view class="user-details">
					<text class="username">{{ userInfo.nickname || userInfo.username }}</text>
					<text class="user-stats">今日已记录 {{ todayRecords }} 项数据</text>
				</view>
			</view>

			<!-- 快捷操作按钮 -->
			<view class="quick-actions">
				<u-button type="primary" icon="plus" @click="addWorkout">记录训练</u-button>
				<u-button type="success" icon="edit-pen" @click="addDiet">记录饮食</u-button>
			</view>
		</view>

		<!-- 数据概览 -->
		<view class="data-overview">
			<u-row gutter="10">
				<u-col span="6">
					<view class="overview-item" @click="toWorkout">
						<u-icon name="level" size="24" color="#3c9cff"></u-icon>
						<text class="item-label">训练记录</text>
						<text class="item-value">{{ workoutCount }}</text>
					</view>
				</u-col>
				<u-col span="6">
					<view class="overview-item" @click="toWeight">
						<u-icon name="weight" size="24" color="#ff9900"></u-icon>
						<text class="item-label">体重</text>
						<text class="item-value">{{ latestWeight }}kg</text>
					</view>
				</u-col>
			</u-row>

			<u-row gutter="10">
				<u-col span="6">
					<view class="overview-item" @click="toMood">
						<u-icon name="heart" size="24" color="#ff6b6b"></u-icon>
						<text class="item-label">心情</text>
						<text class="item-value">{{ moodScore }}/10</text>
					</view>
				</u-col>
				<u-col span="6">
					<view class="overview-item" @click="toDiet">
						<u-icon name="gift" size="24" color="#6bcf7f"></u-icon>
						<text class="item-label">饮食</text>
						<text class="item-value">{{ dietCount }}餐</text>
					</view>
				</u-col>
			</u-row>
		</view>

		<!-- 最近记录 -->
		<view class="recent-records">
			<u-section title="最近记录" :right="true" @click="toAllRecords"></u-section>

			<view class="record-list">
				<view v-for="(record, index) in recentRecords" :key="index" class="record-item">
					<u-icon :name="getRecordIcon(record.type)" :color="getRecordColor(record.type)" size="20"></u-icon>
					<view class="record-content">
						<text class="record-title">{{ record.title }}</text>
						<text class="record-time">{{ record.time }}</text>
					</view>
					<text class="record-value">{{ record.value }}</text>
				</view>
			</view>
		</view>

		<!-- 底部导航栏 -->
		<u-tabbar :value="currentTab" @change="changeTab" fixed placeholder>
			<u-tabbar-item icon="home" text="首页"></u-tabbar-item>
			<u-tabbar-item icon="level" text="训练"></u-tabbar-item>
			<u-tabbar-item icon="heart" text="心情"></u-tabbar-item>
			<u-tabbar-item icon="pie-chart" text="统计"></u-tabbar-item>
			<u-tabbar-item icon="account" text="我的"></u-tabbar-item>
		</u-tabbar>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				currentTab: 0,
				userInfo: {
					username: 'user123',
					nickname: '健身爱好者',
					avatar: ''
				},
				todayRecords: 3,
				workoutCount: 25,
				latestWeight: 70.5,
				moodScore: 8,
				dietCount: 4,
				recentRecords: [
					{ type: 'workout', title: '深蹲训练', time: '今天 18:30', value: '80kg x 3组' },
					{ type: 'diet', title: '午餐记录', time: '今天 12:15', value: '520卡路里' },
					{ type: 'mood', title: '心情打卡', time: '今天 09:00', value: '8分' },
					{ type: 'weight', title: '体重测量', time: '昨天 08:00', value: '70.5kg' }
				]
			}
		},
		methods: {
			changeTab(index) {
				this.currentTab = index
				// 根据tab跳转到不同页面
				const pages = ['/', '/pages/workout/list', '/pages/mood/list', '/pages/statistics/index', '/pages/user/index']
				if (index !== 0) {
					uni.navigateTo({
						url: pages[index]
					})
				}
			},

			addWorkout() {
				uni.navigateTo({
					url: '/pages/workout/add'
				})
			},

			addDiet() {
				uni.navigateTo({
					url: '/pages/diet/add'
				})
			},

			toWorkout() {
				uni.navigateTo({
					url: '/pages/workout/list'
				})
			},

			toWeight() {
				uni.navigateTo({
					url: '/pages/weight/list'
				})
			},

			toMood() {
				uni.navigateTo({
					url: '/pages/mood/list'
				})
			},

			toDiet() {
				uni.navigateTo({
					url: '/pages/diet/list'
				})
			},

			toSetting() {
				uni.navigateTo({
					url: '/pages/user/setting'
				})
			},

			toAllRecords() {
				// 跳转到所有记录页面
			},

			getRecordIcon(type) {
				const icons = {
					'workout': 'level',
					'diet': 'gift',
					'mood': 'heart',
					'weight': 'weight'
				}
				return icons[type] || 'file-text'
			},

			getRecordColor(type) {
				const colors = {
					'workout': '#3c9cff',
					'diet': '#6bcf7f',
					'mood': '#ff6b6b',
					'weight': '#ff9900'
				}
				return colors[type] || '#909399'
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
	padding: 40rpx;
	margin-bottom: 30rpx;
	color: white;

	.user-info {
		display: flex;
		align-items: center;
		margin-bottom: 30rpx;

		.user-details {
			margin-left: 20rpx;

			.username {
				font-size: 36rpx;
				font-weight: bold;
				display: block;
			}

			.user-stats {
				font-size: 24rpx;
				opacity: 0.9;
				margin-top: 10rpx;
				display: block;
			}
		}
	}

	.quick-actions {
		display: flex;
		gap: 20rpx;

		::v-deep .u-btn {
			flex: 1;
		}
	}
}

.data-overview {
	background: white;
	border-radius: 20rpx;
	padding: 30rpx;
	margin-bottom: 30rpx;

	.overview-item {
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
		padding: 30rpx 0;
		text-align: center;

		.item-label {
			font-size: 24rpx;
			color: #666;
			margin: 10rpx 0;
		}

		.item-value {
			font-size: 32rpx;
			font-weight: bold;
			color: #333;
		}
	}
}

.recent-records {
	background: white;
	border-radius: 20rpx;
	padding: 30rpx;

	.record-list {
		margin-top: 20rpx;

		.record-item {
			display: flex;
			align-items: center;
			padding: 20rpx 0;
			border-bottom: 1rpx solid #eee;

			&:last-child {
				border-bottom: none;
			}

			.record-content {
				flex: 1;
				margin: 0 20rpx;

				.record-title {
					font-size: 28rpx;
					color: #333;
					display: block;
				}

				.record-time {
					font-size: 24rpx;
					color: #999;
					margin-top: 5rpx;
					display: block;
				}
			}

			.record-value {
				font-size: 26rpx;
				color: #666;
			}
		}
	}
}
</style>