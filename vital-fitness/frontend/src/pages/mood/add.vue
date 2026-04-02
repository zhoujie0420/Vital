<template>
	<view class="container">
		<!-- 顶部导航栏 -->
		<u-navbar title="记录心情" fixed placeholder>
			<view slot="left">
				<u-icon name="arrow-left" size="20" @click="goBack"></u-icon>
			</view>
			<view slot="right">
				<text class="nav-btn" @click="saveMood">保存</text>
			</view>
		</u-navbar>

		<view class="form-container">
			<!-- 记录日期 -->
			<view class="form-item">
				<text class="label">记录日期</text>
				<u-datetime-picker
					:show="showDatePicker"
					v-model="moodForm.record_date"
					mode="date"
					@confirm="confirmDate"
					@cancel="showDatePicker = false"
				></u-datetime-picker>
				<u-cell
					title="记录日期"
					:value="formatDate(moodForm.record_date)"
					isLink
					@click="showDatePicker = true"
				></u-cell>
			</view>

			<!-- 心情评分 -->
			<view class="form-item">
				<text class="label">心情评分 (1-10分)</text>
				<view class="mood-slider">
					<slider
						v-model="moodForm.mood_score"
						min="1"
						max="10"
						step="1"
						show-value
						activeColor="#ff6b6b"
						block-size="20"
					/>
				</view>

				<!-- 心情表情展示 -->
				<view class="mood-expression">
					<view class="expression-item" v-for="exp in moodExpressions" :key="exp.score"
						:class="{ active: moodForm.mood_score === exp.score }" @click="selectMoodScore(exp.score)">
						<text class="emoji">{{ exp.emoji }}</text>
						<text class="label">{{ exp.label }}</text>
					</view>
				</view>
			</view>

			<!-- 心情标签 -->
			<view class="form-item">
				<text class="label">心情标签</text>
				<view class="tags-container">
					<view class="tag-list">
						<u-tag
							v-for="tag in moodTags"
							:key="tag"
							:text="tag"
							:type="selectedTags.includes(tag) ? 'warning' : 'info'"
							size="mini"
							@click="toggleTag(tag)"
							style="margin: 10rpx;"
						/>
					</view>

					<!-- 自定义标签 -->
					<view class="custom-tag">
						<u-input
							v-model="customTag"
							placeholder="添加自定义标签"
							border="surround"
							size="mini"
							style="flex: 1; margin-right: 20rpx;"
						/>
						<u-button type="primary" size="mini" @click="addCustomTag">添加</u-button>
					</view>
				</view>
			</view>

			<!-- 心情描述 -->
			<view class="form-item">
				<text class="label">心情描述</text>
				<u-textarea
					v-model="moodForm.description"
					placeholder="描述一下今天的心情、发生的事情或感受..."
					:height="200"
				/>
			</view>
		</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				showDatePicker: false,
				customTag: '',

				moodForm: {
					record_date: Date.now(),
					mood_score: 5,
					mood_tags: '',
					description: ''
				},

				selectedTags: [],

				moodTags: ['开心', '兴奋', '满足', '平静', '专注', '疲惫', '焦虑', '沮丧', '愤怒', '孤独', '压力大', '放松'],

				moodExpressions: [
					{ score: 1, emoji: '😫', label: '糟糕' },
					{ score: 2, emoji: '😞', label: '很差' },
					{ score: 3, emoji: '😔', label: '较差' },
					{ score: 4, emoji: '😕', label: '一般' },
					{ score: 5, emoji: '😐', label: '普通' },
					{ score: 6, emoji: '🙂', label: '还好' },
					{ score: 7, emoji: '😊', label: '不错' },
					{ score: 8, emoji: '😄', label: '良好' },
					{ score: 9, emoji: '😁', label: '很好' },
					{ score: 10, emoji: '😍', label: '极佳' }
				]
			}
		},
		watch: {
			selectedTags: {
				handler(newVal) {
					this.moodForm.mood_tags = newVal.join(',')
				},
				deep: true
			}
		},
		methods: {
			goBack() {
				uni.navigateBack()
			},

			saveMood() {
				if (this.moodForm.mood_score < 1 || this.moodForm.mood_score > 10) {
					uni.showToast({
						title: '请选择心情评分',
						icon: 'none'
					})
					return
				}

				// 这里应该调用API保存心情记录
				uni.showToast({
					title: '保存成功',
					icon: 'success'
				})

				// 返回上一页
				setTimeout(() => {
					uni.navigateBack()
				}, 1000)
			},

			confirmDate(e) {
				this.moodForm.record_date = e.value
				this.showDatePicker = false
			},

			selectMoodScore(score) {
				this.moodForm.mood_score = score
			},

			toggleTag(tag) {
				const index = this.selectedTags.indexOf(tag)
				if (index > -1) {
					this.selectedTags.splice(index, 1)
				} else {
					this.selectedTags.push(tag)
				}
			},

			addCustomTag() {
				if (this.customTag.trim() && !this.moodTags.includes(this.customTag.trim())) {
					this.moodTags.push(this.customTag.trim())
					this.toggleTag(this.customTag.trim())
					this.customTag = ''
				}
			},

			formatDate(timestamp) {
				const date = new Date(timestamp)
				return `${date.getFullYear()}-${(date.getMonth() + 1).toString().padStart(2, '0')}-${date.getDate().toString().padStart(2, '0')}`
			}
		}
	}
</script>

<style lang="scss" scoped>
.container {
	padding: 20rpx;
}

.nav-btn {
	color: #3c9cff;
	font-size: 28rpx;
}

.form-container {
	.form-item {
		background: white;
		border-radius: 20rpx;
		padding: 0 30rpx 30rpx;
		margin-bottom: 20rpx;

		.label {
			display: block;
			font-size: 32rpx;
			font-weight: bold;
			color: #333;
			padding: 30rpx 0 20rpx;
		}

		.mood-slider {
			padding: 0 20rpx 30rpx;
		}

		.mood-expression {
			display: flex;
			flex-wrap: wrap;
			justify-content: space-around;

			.expression-item {
				text-align: center;
				padding: 20rpx;
				border-radius: 10rpx;
				margin: 10rpx;
				background: #f5f5f5;

				&.active {
					background: #ff6b6b;
					color: white;
				}

				.emoji {
					font-size: 48rpx;
					display: block;
					margin-bottom: 10rpx;
				}

				.label {
					font-size: 20rpx;
				}
			}
		}

		.tags-container {
			.tag-list {
				margin-bottom: 30rpx;
				min-height: 80rpx;
			}

			.custom-tag {
				display: flex;
				align-items: center;
			}
		}
	}
}
</style>