<template>
	<view class="page" :style="{ paddingTop: topPadding + 'px' }">
		<view class="page-header">
			<text class="nav-back" @tap="goBack">‹</text>
			<text class="page-title">心情记录</text>
			<text class="header-action" @tap="addMood">+ 添加</text>
		</view>

		<view class="list">
			<view v-for="(r, i) in moodList" :key="i" class="mood-card">
				<view class="mood-top">
					<text class="mood-emoji">{{ getEmoji(r.mood_score) }}</text>
					<view class="mood-meta">
						<text class="mood-text">{{ getMoodText(r.mood_score) }}</text>
						<text class="mood-date">{{ formatDate(r.record_date) }}</text>
					</view>
					<view class="mood-score-badge">
						<text class="mood-score">{{ r.mood_score }}/10</text>
					</view>
				</view>
				<view v-if="r.mood_tags" class="mood-tags">
					<text v-for="tag in r.mood_tags.split(',')" :key="tag" class="mood-tag">{{ tag }}</text>
				</view>
				<text v-if="r.description" class="mood-desc">{{ r.description }}</text>
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
@import '../../styles/variables.scss';

.page {
	padding: 0 $spacing-xl;
	padding-bottom: 40rpx;
	min-height: 100vh;
	background: $color-bg;
}

.page-header {
	display: flex;
	justify-content: space-between;
	align-items: center;
	margin-bottom: $spacing-lg;

	.nav-back {
		font-size: 48rpx;
		color: $color-primary;
		font-weight: 300;
		line-height: 1;
		width: 60rpx;
	}
	.page-title {
		font-size: $font-headline;
		font-weight: 600;
		color: $color-label;
	}
	.header-action {
		font-size: $font-subhead;
		color: $color-primary;
		font-weight: 600;
		width: 60rpx;
		text-align: right;
	}
}

// --- Mood Card ---
.mood-card {
	@include card;
	margin-bottom: $spacing-sm;
	@include press-effect;

	.mood-top {
		display: flex;
		align-items: center;

		.mood-emoji {
			font-size: 56rpx;
			margin-right: $spacing-md;
		}

		.mood-meta {
			flex: 1;

			.mood-text {
				display: block;
				font-size: $font-headline;
				font-weight: 600;
				color: $color-label;
			}
			.mood-date {
				display: block;
				font-size: $font-caption1;
				color: $color-label-quaternary;
				margin-top: 2rpx;
			}
		}

		.mood-score-badge {
			background: $color-fill;
			padding: $spacing-xs $spacing-md;
			border-radius: $spacing-xs;

			.mood-score {
				font-size: $font-subhead;
				font-weight: 700;
				color: $color-label-quaternary;
				font-variant-numeric: tabular-nums;
			}
		}
	}

	.mood-tags {
		display: flex;
		flex-wrap: wrap;
		gap: $spacing-xs;
		margin-top: $spacing-md;

		.mood-tag {
			padding: 6rpx $spacing-md;
			background: $color-primary-light;
			color: $color-primary;
			border-radius: $radius-full;
			font-size: $font-caption1;
			font-weight: 500;
		}
	}

	.mood-desc {
		display: block;
		font-size: $font-subhead;
		color: $color-label-tertiary;
		margin-top: $spacing-md;
		line-height: 1.6;
	}
}

.empty {
	@include empty-state;
}
</style>
