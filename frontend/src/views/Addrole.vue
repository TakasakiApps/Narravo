<template>
	<div style="margin: 0; padding: 0;">
		<el-carousel height="93vh" direction="horizontal" :autoplay="false" arrow="never" indicator-position="none"
			ref="carousel">
			<el-carousel-item :key="1" class="mian">
				<div v-if="true">
					<h1 class="mian-text">
						暂无角色
						<div>
							请点击左下方号以继续
						</div>
					</h1>
				</div>
				<el-affix position="bottom" :offset="50" class="fixed-element" @click="addRoleCard" v-show="button[0]">
					<svg height="50" viewBox="-30 -750 600 900" width="50" class="btns addCard">
						<g transform="scale(0.6)">
							<path d="M450-200v-250H200v-60h250v-250h60v250h250v60H510v250h-60Z" fill="currentColor" />
						</g>
					</svg>
				</el-affix>
			</el-carousel-item>
			<el-carousel-item :key="2">
				<div>
					<el-card class="box-card">
						<div class="guide">
							<div class="guide-child">
								创建您的角色
							</div>
							<div class="guide-child">
								作为小说朗读者
							</div>
						</div>
						<el-card class="box-card-main">
							<el-carousel height="65vh" direction="vertical" :autoplay="false" arrow="never">
								<el-carousel-item :key="2 - 1">

									<el-avatar :size="200" class="avatar" />
									<div class="Roleinput">
										<el-input v-model="Card.name" placeholder="请输入你想要创建的角色名" />
										<div class="avatar-txt">
											通过创建自定义角色来个性化有声小说
											<div>跟随向导以创建您的自定义角色</div>
										</div>
									</div>
								</el-carousel-item>
								<el-carousel-item :key="2 - 2">
									<div class="cardValue">
										<div class="cardValue-title">语言模型</div>
										<el-select v-model="Card.Reader" class="cardInput" placeholder="Select" size="default">
											<el-option label="成男1号姬" value="成男1" />
										</el-select>
									</div>
									<div>
										<div class="cardValue-title">默认情感</div>
										<el-select v-model="Card" class="cardInput" placeholder="Select" size="default">
											<el-option label="不支持" value="no" />
										</el-select>
									</div>
									<div>
										<div class="cardValue-title">情感等级</div>
										<el-slider v-model="Card.emotion" class="cardInput" :step="1" size="small" show-stops :max="200" />
									</div>
									<div>
										<div class="cardValue-title">语速</div>
										<el-slider v-model="Card.speed" :step="0.5" class="cardInput" size="small" show-stops :max="2.5"
											:min="-2" />
									</div>
									<div>
										<div class="cardValue-title">音量</div>
										<el-slider v-model="Card.tone" :step="1" class="cardInput" size="small" show-stops :max="10" />
									</div>
								</el-carousel-item>
							</el-carousel>
						</el-card>
					</el-card>

				</div>
			</el-carousel-item>
		</el-carousel>
	</div>
</template>
    
<script setup lang='ts'>
import RoleCard from '../components/roleCards/RoleCard.vue';
import readRoledir from '../components/roleCards/node-api'
import { onMounted, reactive, ref } from 'vue';
import { StarFilled, Plus } from '@element-plus/icons-vue'
onMounted(() => {
	// const e = ReadDir()
	// if (Array.isArray(e)) {
	// 	// 对数组类型进行断言
	// 	// 使用 s.length 进行进一步的验证

	// 	if (e.length === 0) {
	// 		// console.log('没有角色卡');
	// 	} else {
	// 		// console.log('具有角色卡');
	// 		console.log(e)
	// 	}
	// } else {
	// 	// 对其他类型进行处理（例如字符串或布尔值）
	// 	console.log('出现意料之外的错误')
	// }
})
const Card = reactive({
	name: '',
	rgb: '111',
	Reader: '小帅',
	speed: 0,
	tone: 5,
	emotion: 100
})
const button = reactive([
	true,
	false,
])
const carousel = ref<HTMLFormElement>()
function buttonShow() {
	button[0] = false;
	button[1] = true
}

function addRoleCard() {
	if (carousel.value) {
		setTimeout(buttonShow, 500)
		carousel.value.next()
	}

}
</script>
    
<style scoped>
.addCard {
	color: white;
	border-radius: 100px;
	box-shadow: 1px 1px 3px rgba(71, 56, 56, 0.661);
	/* margin-right: 2px; */
}

.fixed-element {
	/* position: absolute;
	bottom: 20px;
	right: 20px; */
	/* display: flex;
	float: left; */
	margin-top: auto;
	margin-left: auto;
	margin-right: 5px;
	/* transform: translateX(50%); */
}

.btns {}

.mian-text {
	position: relative;
	right: 35vw;
	top: 45vh;
	text-align: center;
	color: rgb(136, 134, 134);
}

.mian {
	display: flex;
	flex-direction: column;
	justify-content: flex-end;
	align-items: flex-end;
}

.box-card {
	position: absolute;
	width: 70vw;
	height: 71vh;
	left: 10%;
	margin-top: 50px;
	background-image: linear-gradient(to right top, #3a586b, #405d7a, #4d6288, #606593, #77669b, #93699e, #ad6c9c, #c57097, #dd7d8e, #ed8f85, #f6a47e, #f6bc7d);
}

.box-card-main {
	width: 342px;
	height: 65vh;
	/* margin-left: 50vh; */
	position: absolute;
	top: 5%;
	right: 20px;

}

.guide {
	position: absolute;
	left: 15%;
	top: 40%;
	font-size: 3.2vw;
	font-weight: 900;
	text-align: center;
	background-image: -webkit-linear-gradient(right top, #cdd1d3, #c9d3d8, #c5d6dc, #c0d8df, #bbdbe2, #bbdbe2, #bbdbe2, #bbdbe2, #c0d8df, #c5d6dc, #c9d3d8, #cdd1d3);
	-webkit-background-clip: text;
	-webkit-text-fill-color: transparent;
}

.guide-child {
	display: block;
}

.avatar {
	position: absolute;
	left: 20%;
	/* top: 50px; */
}

.Roleinput {
	position: absolute;
	bottom: 75px;
	left: 5%;
}

.avatar-txt {

	text-align: center;

	color: #949494;
}

.cardInput {
	/* display: inline; */
	width: 20vw;
}

.cardValue-title {
	font-size: 12px;
}

.cardInput {}
</style>