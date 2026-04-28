<template>
	<view class="page" :style="{ paddingTop: topPadding + 'px' }">
		<view class="nav-bar">
			<text class="nav-back" @tap="goBack">‹ 返回</text>
			<text class="nav-title">饮食规划</text>
			<view class="nav-placeholder"></view>
		</view>

		<!-- 当前方案 — 压缩为一行 -->
		<view v-if="store.activePlan" class="active-bar" @tap="editPlan(store.activePlan.id)">
			<view class="ab-left">
				<view class="ab-dot"></view>
				<text class="ab-name">{{ store.activePlan.name }}</text>
				<text class="ab-cal">{{ store.activePlan.macroTargets.calories }} kcal</text>
			</view>
			<text class="ab-chevron">›</text>
		</view>
		<view v-else class="no-plan-bar">
			<text class="np-text">暂未启用方案</text>
		</view>

		<!-- 我的方案列表 -->
		<view v-if="store.plans.length > 0" class="section">
			<text class="section-title">我的方案</text>
			<view v-for="plan in store.plans" :key="plan.id" class="plan-row">
				<view class="plan-row-left" @tap="editPlan(plan.id)">
					<text class="pr-name">{{ plan.name }}</text>
					<text class="pr-detail">{{ typeLabel(plan.type) }} · {{ plan.macroTargets.calories }}kcal · 蛋白{{ plan.macroTargets.protein }}g</text>
				</view>
				<view class="plan-row-right">
					<view v-if="plan.isActive" class="pr-badge active"><text>使用中</text></view>
					<text v-else class="pr-action" @tap="activate(plan.id)">启用</text>
					<text class="pr-delete" @tap="deletePlan(plan)">删除</text>
				</view>
			</view>
		</view>

		<!-- 快速创建 — 直接展示模板 -->
		<view class="section">
			<text class="section-title">从模板创建</text>
			<view v-for="tpl in systemTemplates" :key="tpl.id" class="tpl-row" @tap="createFromTemplate(tpl.id)">
				<view class="tpl-row-left">
					<text class="tr-name">{{ tpl.name }}</text>
					<text class="tr-desc">{{ tpl.description }}</text>
				</view>
				<text class="tr-chevron">›</text>
			</view>
		</view>

		<!-- 自定义 -->
		<view class="custom-btn" @tap="createCustom">
			<text>+ 自定义方案</text>
		</view>
	</view>
</template>

<script setup>
	import { computed } from 'vue'
	import { onShow } from '@dcloudio/uni-app'
	import { useDietPlanStore } from '../../store/dietPlan'
	import { PUBLIC_TEMPLATES } from '../../utils/dietPlanTemplates'

	const store = useDietPlanStore()
	const systemTemplates = PUBLIC_TEMPLATES

	const topPadding = computed(() => {
		const app = getApp()
		return (app.globalData?.customBarHeight || 88) + 8
	})

	onShow(() => store.init())

	const goBack = () => uni.navigateBack()

	const typeLabel = (type) => {
		return { fixed: '固定比例', carb_cycle: '碳循环', carb_taper: '碳水渐降' }[type] || type
	}

	const activate = (id) => {
		store.activatePlan(id)
		uni.showToast({ title: '已启用', icon: 'success' })
	}

	const editPlan = (id) => {
		uni.navigateTo({ url: '/pages/diet/plan-edit?mode=edit&planId=' + id })
	}

	const deletePlan = (plan) => {
		uni.showModal({
			title: '删除方案',
			content: `确定删除「${plan.name}」？`,
			success: (res) => {
				if (res.confirm) {
					store.deletePlan(plan.id)
					uni.showToast({ title: '已删除', icon: 'success' })
				}
			}
		})
	}

	// 点模板直接进编辑页，一步到位
	const createFromTemplate = (templateId) => {
		uni.navigateTo({ url: '/pages/diet/plan-edit?mode=template&templateId=' + templateId })
	}

	const createCustom = () => {
		uni.navigateTo({ url: '/pages/diet/plan-edit?mode=custom' })
	}
</script>

<style lang="scss" scoped>
@import '../../styles/variables.scss';

.page { padding: 0 $spacing-xl; padding-bottom: 40rpx; min-height: 100vh; background: $color-bg; }

.nav-bar { @include nav-bar;
	.nav-back { font-size: $font-callout; color: $color-primary; font-weight: 500; }
	.nav-title { font-size: $font-headline; font-weight: 600; color: $color-label; }
	.nav-placeholder { width: 80rpx; }
}

// --- Active bar (compact) ---
.active-bar {
	@include card;
	padding: $spacing-md $spacing-xl;
	margin-bottom: $spacing-lg;
	display: flex;
	justify-content: space-between;
	align-items: center;
	background: $color-label;
	@include press-effect;

	.ab-left { display: flex; align-items: center; gap: $spacing-sm; }
	.ab-dot { width: 12rpx; height: 12rpx; border-radius: 50%; background: $color-primary; }
	.ab-name { font-size: $font-body; font-weight: 600; color: #fff; }
	.ab-cal { font-size: $font-caption1; color: rgba(255,255,255,0.5); font-variant-numeric: tabular-nums; }
	.ab-chevron { font-size: 36rpx; color: rgba(255,255,255,0.3); font-weight: 300; }
}

.no-plan-bar {
	@include card;
	padding: $spacing-lg $spacing-xl;
	margin-bottom: $spacing-lg;
	text-align: center;

	.np-text { font-size: $font-subhead; color: $color-label-quaternary; }
}

// --- Section ---
.section { margin-bottom: $spacing-lg; }
.section-title {
	display: block;
	font-size: $font-footnote;
	font-weight: 600;
	color: $color-label-quaternary;
	text-transform: uppercase;
	letter-spacing: 2rpx;
	margin-bottom: $spacing-md;
}

// --- Plan rows ---
.plan-row {
	@include card;
	padding: $spacing-md $spacing-xl;
	margin-bottom: $spacing-xs;
	display: flex;
	justify-content: space-between;
	align-items: center;

	.plan-row-left {
		flex: 1;
		@include press-effect;

		.pr-name { display: block; font-size: $font-body; font-weight: 600; color: $color-label; }
		.pr-detail { display: block; font-size: $font-caption2; color: $color-label-quaternary; margin-top: 2rpx; font-variant-numeric: tabular-nums; }
	}

	.plan-row-right {
		display: flex;
		align-items: center;
		gap: $spacing-md;
		flex-shrink: 0;

		.pr-badge {
			padding: 4rpx $spacing-sm;
			border-radius: $spacing-xs;
			font-size: $font-caption2;
			font-weight: 600;

			&.active { background: $color-primary-light; color: $color-primary; }
		}
		.pr-action { font-size: $font-caption1; color: $color-primary; font-weight: 600; &:active { opacity: 0.5; } }
		.pr-delete { font-size: $font-caption1; color: $color-red; font-weight: 500; &:active { opacity: 0.5; } }
	}
}

// --- Template rows (direct, no intermediate page) ---
.tpl-row {
	@include card;
	padding: $spacing-md $spacing-xl;
	margin-bottom: $spacing-xs;
	display: flex;
	justify-content: space-between;
	align-items: center;
	@include press-effect;

	.tpl-row-left {
		flex: 1;
		.tr-name { display: block; font-size: $font-body; font-weight: 600; color: $color-label; }
		.tr-desc { display: block; font-size: $font-caption2; color: $color-label-quaternary; margin-top: 2rpx; line-height: 1.4; }
	}
	.tr-chevron { font-size: 36rpx; color: $color-separator-opaque; font-weight: 300; flex-shrink: 0; }
}

// --- Custom button ---
.custom-btn {
	@include card;
	text-align: center;
	padding: $spacing-lg;
	margin-top: $spacing-sm;
	@include press-effect;

	text { font-size: $font-subhead; color: $color-primary; font-weight: 600; }
}
</style>
