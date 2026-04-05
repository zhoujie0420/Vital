<template>
	<view class="container">
		<u-navbar title="心情记录" fixed placeholder>
			<view slot="right">
				<text class="add-btn" @tap="addMood">+ 记录</text>
			</view>
		</u-navbar>

		<view class="mood-list">
			<view v-for="(r, i) in moodList" :key="i" class="mood-item">
				<view class="item-header">
					<text class="item-date">{{ formatDate(r.record_date) }}</text>
					<text class="item-score">{{ r.mood_score }}/10</text>
				</view>
				<view v-if="r.mood_tags" class="item-tags">
					<text v-for="tag in r.mood_tags.split(',')" :key="tag" class="tag">{{ tag }}</text>
				</view>
				<text v-if="r.description" class="item-desc">{{ r.description }}</text>
			</view>

			<view v-if="moodList.length === 0" class="empty-state">
				<text class="empty-text">暂无心情记录</text>
				<button class="add-first-btn" type="primary" @tap="addMood">记录今天的心情</button>
			</view>
		</view>
	</view>
</template>

<script>
	import { getMoods } from '../../api/mood'

	export default {
		data() { return { moodList: [] } },
		onShow() { this.load() },
		methods: {
			async load() {
				try {
					const res = await getMoods({ page: 1, page_size: 50 })
					this.moodList = res.data || []
				} catch (e) {}
			},
			addMood() { uni.navigateTo({ url: '/pages/mood/add' }) },
			formatDate(d) {
				if (!d) return ''
				const date = new Date(d)
				return (date.getMonth()+1) + '月' + date.getDate() + '日'
			}
		}
	}
</script>

<style lang="scss" scoped>
.container { padding: 20rpx; }
.add-btn { color: #3c9cff; font-size: 28rpx; font-weight: bold; }
.mood-list {
	.mood-item {
		background: white; border-radius: 20rpx; padding: 30rpx; margin-bottom: 15rpx;
		.item-header { display: flex; justify-content: space-between; margin-bottom: 10rpx;
			.item-date { font-size: 24rpx; color: #999; }
			.item-score { font-size: 32rpx; font-weight: bold; color: #ff6b6b; }
		}
		.item-tags { display: flex; flex-wrap: wrap; gap: 10rpx; margin-bottom: 10rpx;
			.tag { padding: 5rpx 15rpx; background: #fff0f0; color: #ff6b6b; border-radius: 10rpx; font-size: 22rpx; }
		}
		.item-desc { font-size: 26rpx; color: #666; line-height: 1.5; }
	}
	.empty-state { text-align: center; padding: 100rpx 0; background: white; border-radius: 20rpx;
		.empty-text { font-size: 28rpx; color: #999; display: block; margin-bottom: 30rpx; }
		.add-first-btn { width: 60%; font-size: 28rpx; }
	}
}
</style>
