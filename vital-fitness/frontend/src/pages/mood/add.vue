<template>
	<view class="page" :style="{ paddingTop: topPadding + 'px' }">
		<view class="nav-bar">
			<text class="nav-back" @tap="goBack">‹ 返回</text>
			<text class="nav-title">记录心情</text>
			<text class="nav-placeholder"></text>
		</view>

		<view class="card">
			<text class="card-label">今天心情怎么样？</text>
			<view class="mood-grid">
				<view v-for="exp in moods" :key="exp.score" class="mood-card"
					:class="{ selected: score === exp.score }" @tap="score = exp.score">
					<text class="emoji">{{ exp.emoji }}</text>
					<text class="mood-text">{{ exp.label }}</text>
				</view>
			</view>
		</view>

		<view class="card" v-if="score > 0">
			<text class="card-label">标签</text>
			<view class="tag-grid">
				<text v-for="tag in tags" :key="tag" class="tag-pill"
					:class="{ active: selectedTags.includes(tag) }" @tap="toggleTag(tag)">
					{{ tag }}
				</text>
			</view>
		</view>

		<view class="card" v-if="score > 0">
			<text class="card-label">说点什么</text>
			<textarea v-model="description" placeholder="记录今天的心情..." class="desc-area" />
		</view>

		<view class="bottom-action" v-if="score > 0">
			<view class="save-btn" :class="{ loading: saving }" @tap="save">
				<text>{{ saving ? '保存中...' : '保存记录' }}</text>
			</view>
		</view>
	</view>
</template>

<script>
	import { createMood } from '../../api/mood'

	export default {
		data() {
			return {
				saving: false, score: 0, description: '', selectedTags: [],
				moods: [
					{ score: 1, emoji: '😫', label: '糟糕' },
					{ score: 3, emoji: '😔', label: '较差' },
					{ score: 5, emoji: '😐', label: '一般' },
					{ score: 7, emoji: '😊', label: '不错' },
					{ score: 9, emoji: '😁', label: '很好' },
					{ score: 10, emoji: '😍', label: '极佳' }
				],
				tags: ['开心', '兴奋', '平静', '专注', '疲惫', '焦虑', '压力大', '放松']
			}
		},
		computed: {
			topPadding() {
				const app = getApp()
				return (app.globalData?.customBarHeight || 88) + 8
			}
		},
		methods: {
			goBack() { uni.navigateBack() },
			toggleTag(tag) {
				const i = this.selectedTags.indexOf(tag)
				if (i > -1) this.selectedTags.splice(i, 1)
				else this.selectedTags.push(tag)
			},
			async save() {
				this.saving = true
				try {
					await createMood({
						record_date: new Date().toISOString(),
						mood_score: this.score,
						mood_tags: this.selectedTags.join(','),
						description: this.description
					})
					uni.showToast({ title: '保存成功', icon: 'success' })
					setTimeout(() => uni.navigateBack(), 1000)
				} catch (e) {} finally { this.saving = false }
			}
		}
	}
</script>

<style lang="scss" scoped>
.page {
	padding: 0 32rpx;
	padding-bottom: 160rpx;
	min-height: 100vh;
	background: #f2f2f7;
}

.nav-bar {
	display: flex;
	align-items: center;
	justify-content: space-between;
	margin-bottom: 28rpx;

	.nav-back { font-size: 32rpx; color: #007aff; font-weight: 500; }
	.nav-title { font-size: 34rpx; font-weight: 600; color: #1c1c1e; }
	.nav-placeholder { width: 80rpx; }
}

.card {
	background: #fff;
	border-radius: 20rpx;
	padding: 28rpx 32rpx;
	margin-bottom: 16rpx;
	box-shadow: 0 2rpx 16rpx rgba(0, 0, 0, 0.04);

	.card-label {
		display: block;
		font-size: 26rpx;
		font-weight: 600;
		color: #8e8e93;
		text-transform: uppercase;
		letter-spacing: 1rpx;
		margin-bottom: 20rpx;
	}
}

.mood-grid {
	display: flex;
	flex-wrap: wrap;
	gap: 16rpx;
	justify-content: center;

	.mood-card {
		width: 160rpx;
		text-align: center;
		padding: 28rpx 12rpx;
		border-radius: 24rpx;
		background: #f2f2f7;
		border: 3rpx solid transparent;
		transition: all 0.25s cubic-bezier(0.25, 0.46, 0.45, 0.94);

		&.selected {
			background: #fff;
			border-color: #007aff;
			box-shadow: 0 6rpx 24rpx rgba(0, 122, 255, 0.2);
			transform: scale(1.05);
		}
		&:active { transform: scale(0.92); }

		.emoji {
			display: block;
			font-size: 64rpx;
			margin-bottom: 10rpx;
		}
		.mood-text {
			display: block;
			font-size: 24rpx;
			color: #636366;
			font-weight: 600;
		}
	}
}

.tag-grid {
	display: flex;
	flex-wrap: wrap;
	gap: 12rpx;

	.tag-pill {
		padding: 14rpx 28rpx;
		border-radius: 24rpx;
		background: #f2f2f7;
		font-size: 26rpx;
		color: #636366;
		font-weight: 500;
		transition: all 0.2s ease;

		&.active {
			background: #007aff;
			color: #fff;
			box-shadow: 0 4rpx 12rpx rgba(0, 122, 255, 0.3);
		}
		&:active { transform: scale(0.92); }
	}
}

.desc-area {
	width: 100%;
	height: 160rpx;
	font-size: 30rpx;
	color: #1c1c1e;
	line-height: 1.6;
}

.bottom-action {
	position: fixed;
	bottom: 0;
	left: 0;
	right: 0;
	padding: 20rpx 32rpx;
	padding-bottom: calc(20rpx + env(safe-area-inset-bottom));
	background: rgba(242, 242, 247, 0.85);
	backdrop-filter: blur(20px);
	-webkit-backdrop-filter: blur(20px);

	.save-btn {
		background: #007aff;
		color: #fff;
		text-align: center;
		padding: 28rpx 0;
		border-radius: 16rpx;
		font-size: 32rpx;
		font-weight: 600;
		box-shadow: 0 4rpx 16rpx rgba(0, 122, 255, 0.3);
		transition: all 0.15s ease;

		&:active { transform: scale(0.98); opacity: 0.9; }
		&.loading { opacity: 0.6; }
	}
}
</style>
