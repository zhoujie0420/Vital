<template>
	<view class="page" :style="{ paddingTop: topPadding + 'px' }">
		<view class="nav-bar">
			<text class="nav-back" @tap="goBack">‹ 返回</text>
			<text class="nav-title">记录心情</text>
			<view class="nav-placeholder"></view>
		</view>

		<view class="card">
			<text class="card-label">今天心情怎么样？</text>
			<view class="mood-grid">
				<view v-for="exp in moods" :key="exp.score" class="mood-item"
					:class="{ selected: score === exp.score }" @tap="score = exp.score">
					<view class="mood-face" :style="{ background: exp.color }">
						<text class="mood-score-num">{{ exp.score }}</text>
					</view>
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
					{ score: 1, icon: ':(', label: '糟糕', color: '#EF4444' },
					{ score: 3, icon: ':|', label: '较差', color: '#F97316' },
					{ score: 5, icon: '-', label: '一般', color: '#EAB308' },
					{ score: 7, icon: ':)', label: '不错', color: '#10B981' },
					{ score: 9, icon: ':D', label: '很好', color: '#14B8A6' },
					{ score: 10, icon: '<3', label: '极佳', color: '#EC4899' }
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
@import '../../styles/variables.scss';

.page {
	padding: 0 $spacing-xl;
	padding-bottom: 160rpx;
	min-height: 100vh;
	background: $color-bg;
}

.nav-bar {
	@include nav-bar;
	.nav-back { font-size: $font-callout; color: $color-primary; font-weight: 500; }
	.nav-title { font-size: $font-headline; font-weight: 600; color: $color-label; }
	.nav-placeholder { width: 80rpx; }
}

.card {
	@include card;
	margin-bottom: $spacing-md;
	.card-label { @include card-label; }
}

// --- Mood Grid ---
.mood-grid {
	display: flex;
	flex-wrap: wrap;
	gap: $spacing-sm;
	justify-content: center;

	.mood-item {
		width: 152rpx;
		text-align: center;
		padding: $spacing-lg $spacing-sm;
		border-radius: $radius-xl;
		background: $color-fill;
		border: 2.5rpx solid transparent;
		transition: all 0.3s cubic-bezier(0.16, 1, 0.3, 1);

		&.selected {
			background: $color-bg-elevated;
			border-color: $color-primary;
			box-shadow: 0 4rpx 20rpx rgba(16, 185, 129, 0.15);
			transform: scale(1.05);
		}
		&:active { transform: scale(0.92); }

		.mood-face {
			width: 64rpx;
			height: 64rpx;
			border-radius: 50%;
			display: flex;
			align-items: center;
			justify-content: center;
			margin: 0 auto $spacing-sm;
			opacity: 0.85;
		}
		.mood-score-num {
			font-size: $font-subhead;
			font-weight: 700;
			color: #fff;
		}
		.mood-text {
			display: block;
			font-size: $font-caption1;
			color: $color-label-tertiary;
			font-weight: 600;
		}
	}
}

// --- Tags ---
.tag-grid {
	display: flex;
	flex-wrap: wrap;
	gap: $spacing-sm;

	.tag-pill {
		padding: $spacing-sm $spacing-lg;
		border-radius: $radius-full;
		background: $color-fill;
		font-size: $font-footnote;
		color: $color-label-tertiary;
		font-weight: 500;
		transition: all 0.25s cubic-bezier(0.25, 0.46, 0.45, 0.94);

		&.active {
			background: $color-label;
			color: #fff;
		}
		&:active { transform: scale(0.92); }
	}
}

.desc-area {
	width: 100%;
	height: 160rpx;
	font-size: $font-body;
	color: $color-label;
	line-height: 1.6;
}

.bottom-action {
	@include bottom-action-bar;
	.save-btn { @include primary-button; }
}
</style>
