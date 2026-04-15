<template>
	<view class="page" :style="{ paddingTop: topPadding + 'px' }">
		<view class="page-header">
			<text class="nav-back" @tap="goBack">‹</text>
			<text class="page-title">心情记录</text>
			<text class="header-action" @tap="addMood">+ 添加</text>
		</view>

		<view class="list">
			<view v-for="(r, i) in moodList" :key="i" class="list-item">
				<view class="item-top">
					<text class="item-emoji">{{ getEmoji(r.mood_score) }}</text>
					<view class="item-meta">
						<text class="item-mood-text">{{ getMoodText(r.mood_score) }}</text>
						<text class="item-date">{{ formatDate(r.record_date) }}</text>
					</view>
					<text class="item-score">{{ r.mood_score }}/10</text>
				</view>
				<view v-if="r.mood_tags" class="item-tags">
					<text v-for="tag in r.mood_tags.split(',')" :key="tag" class="tag">{{ tag }}</text>
				</view>
				<text v-if="r.description" class="item-desc">{{ r.description }}</text>
			</view>
		</view>

		<view v-if="moodList.length === 0" class="empty">
			<text class="empty-icon">😊</text>
			<text class="empty-title">暂无心情记录</text>
			<text class="empty-desc">记录每天的心情，关注内心变化</text>
			<view class="empty-btn" @tap="addMood">
				<text>开始记录</text>
			</view>
		</view>
	</view>
</template>

<script>
	import { getMoods } from '../../api/mood'

	export default {
		data() { return { moodList: [] } },
		computed: {
			topPadding() {
				const app = getApp()
				return (app.globalData?.customBarHeight || 88) + 8
			}
		},
		onShow() { this.load() },
		methods: {
			async load() {
				try {
					const res = await getMoods({ page: 1, page_size: 50 })
					this.moodList = res.data || []
				} catch (e) {}
			},
			addMood() { uni.navigateTo({ url: '/pages/mood/add' }) },
			goBack() { uni.navigateBack() },
			formatDate(d) {
				if (!d) return ''
				const date = new Date(d)
				return (date.getMonth()+1) + '月' + date.getDate() + '日'
			},
			getEmoji(score) {
				if (score >= 9) return '😍'
				if (score >= 7) return '😊'
				if (score >= 5) return '😐'
				if (score >= 3) return '😔'
				return '😫'
			},
			getMoodText(score) {
				if (score >= 9) return '极佳'
				if (score >= 7) return '不错'
				if (score >= 5) return '一般'
				if (score >= 3) return '较差'
				return '糟糕'
			}
		}
	}
</script>

<style lang="scss" scoped>
.page {
	padding: 0 32rpx;
	padding-bottom: 40rpx;
	min-height: 100vh;
	background: #f2f2f7;
}

.page-header {
	display: flex;
	justify-content: space-between;
	align-items: center;
	margin-bottom: 28rpx;

	.nav-back {
		font-size: 48rpx;
		color: #007aff;
		font-weight: 300;
		line-height: 1;
	}

	.page-title {
		font-size: 34rpx;
		font-weight: 600;
		color: #1c1c1e;
	}
	.header-action {
		font-size: 30rpx;
		color: #007aff;
		font-weight: 600;
	}
}

.list-item {
	background: #fff;
	border-radius: 20rpx;
	padding: 28rpx 32rpx;
	margin-bottom: 16rpx;
	box-shadow: 0 2rpx 16rpx rgba(0, 0, 0, 0.04);
	transition: transform 0.15s ease;

	&:active { transform: scale(0.98); }

	.item-top {
		display: flex;
		align-items: center;

		.item-emoji {
			font-size: 64rpx;
			margin-right: 20rpx;
		}

		.item-meta {
			flex: 1;

			.item-mood-text {
				display: block;
				font-size: 32rpx;
				font-weight: 600;
				color: #1c1c1e;
			}
			.item-date {
				display: block;
				font-size: 24rpx;
				color: #8e8e93;
				margin-top: 4rpx;
			}
		}

		.item-score {
			font-size: 28rpx;
			font-weight: 700;
			color: #8e8e93;
			background: #f2f2f7;
			padding: 8rpx 16rpx;
			border-radius: 12rpx;
		}
	}

	.item-tags {
		display: flex;
		flex-wrap: wrap;
		gap: 12rpx;
		margin-top: 20rpx;

		.tag {
			padding: 8rpx 20rpx;
			background: rgba(0, 122, 255, 0.08);
			color: #007aff;
			border-radius: 20rpx;
			font-size: 24rpx;
			font-weight: 500;
		}
	}

	.item-desc {
		display: block;
		font-size: 28rpx;
		color: #636366;
		margin-top: 16rpx;
		line-height: 1.6;
	}
}

.empty {
	text-align: center;
	padding: 120rpx 0;

	.empty-icon { display: block; font-size: 96rpx; margin-bottom: 24rpx; }
	.empty-title { display: block; font-size: 34rpx; font-weight: 600; color: #1c1c1e; margin-bottom: 8rpx; }
	.empty-desc { display: block; font-size: 26rpx; color: #8e8e93; margin-bottom: 40rpx; }
	.empty-btn {
		display: inline-block;
		padding: 20rpx 56rpx;
		background: #007aff;
		color: #fff;
		border-radius: 16rpx;
		font-size: 30rpx;
		font-weight: 600;
		box-shadow: 0 4rpx 16rpx rgba(0, 122, 255, 0.3);

		&:active { transform: scale(0.95); opacity: 0.9; }
	}
}
</style>
