<template>
	<view class="container">
		<u-navbar title="记录心情" fixed placeholder>
			<view slot="left">
				<u-icon name="arrow-left" size="20" @click="goBack"></u-icon>
			</view>
		</u-navbar>

		<view class="form-container">
			<!-- 心情评分 -->
			<view class="form-item">
				<text class="label">今天心情怎么样？</text>
				<view class="mood-grid">
					<view
						v-for="exp in moods"
						:key="exp.score"
						class="mood-card"
						:class="{ selected: score === exp.score }"
						@tap="score = exp.score"
					>
						<text class="emoji">{{ exp.emoji }}</text>
						<text class="mood-label">{{ exp.label }}</text>
					</view>
				</view>
			</view>

			<!-- 心情标签 -->
			<view class="form-item" v-if="score > 0">
				<text class="label">标签</text>
				<view class="tag-list">
					<text
						v-for="tag in tags"
						:key="tag"
						class="tag-item"
						:class="{ active: selectedTags.includes(tag) }"
						@tap="toggleTag(tag)"
					>{{ tag }}</text>
				</view>
			</view>

			<!-- 描述 -->
			<view class="form-item" v-if="score > 0">
				<text class="label">说点什么</text>
				<textarea v-model="description" placeholder="记录今天的心情..." class="desc-input" />
			</view>

			<button v-if="score > 0" class="save-btn" type="primary" :loading="saving" @tap="save">
				{{ saving ? '保存中...' : '保存记录' }}
			</button>
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
.container { padding: 20rpx; }
.form-container {
	.form-item {
		background: white; border-radius: 20rpx; padding: 30rpx; margin-bottom: 20rpx;
		.label { font-size: 32rpx; font-weight: bold; color: #333; margin-bottom: 20rpx; display: block; }
		.mood-grid {
			display: flex; flex-wrap: wrap; gap: 15rpx; justify-content: center;
			.mood-card {
				width: 140rpx; text-align: center; padding: 20rpx 10rpx; border-radius: 15rpx; background: #f5f5f5;
				border: 2rpx solid transparent;
				&.selected { background: #fff0f0; border-color: #ff6b6b; }
				.emoji { font-size: 48rpx; display: block; margin-bottom: 8rpx; }
				.mood-label { font-size: 22rpx; color: #666; }
			}
		}
		.tag-list {
			display: flex; flex-wrap: wrap; gap: 15rpx;
			.tag-item {
				padding: 10rpx 24rpx; border-radius: 30rpx; background: #f5f5f5; font-size: 24rpx; color: #666;
				&.active { background: #ff6b6b; color: white; }
			}
		}
		.desc-input { width: 100%; height: 150rpx; border: 1rpx solid #eee; border-radius: 10rpx; padding: 20rpx; font-size: 28rpx; }
	}
	.save-btn { margin-top: 20rpx; border-radius: 15rpx; font-size: 32rpx; }
}
</style>
