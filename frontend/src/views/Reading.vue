<template>
	<div class="container openscr" ref="scrool">
		<h1 style="display: inline-block;">{{ optison[progressIndex].name }}</h1>
		<div class="Progress">
			{{ percentage }}
		</div>
		<div>
			<p class="dom" :style="{ '--font-size': fontSize + 'px' }" v-for="(item, index) in progrdata" :key="index">{{ item
			}}
			</p>
		</div>
		<div class="sidenav">
			<el-aside :class="[right == false ? 'open' : '']" id="aside">
				<div class="body">
					<h2>{{ removeExtension(props.name) }}</h2>
					<el-scrollbar max-height="480px" always="true">
						<div class="chapter" v-for="(item, index) in optison" :key="index" :cid="item.chapterId"
							@click="changeper(item, index)" id="chapter-ilem">
							<el-text truncated>{{ item.name }}</el-text>
						</div>
					</el-scrollbar>
				</div>

				<div class="bottom">
					<p style="font-family: '更纱黑体 SC'; font-size: 14px;">字体大小</p>
					<el-input-number class="el-input-mian" v-model="fontSize"
						style="width: 78px; border-radius: 25% ;  margin-left: 12px; display: flex; justify-content: center; color: brown;"></el-input-number>
				</div>
			</el-aside>
			<div class="right" @click="checkLight">
				<svg t="1688729279262" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg"
					p-id="2586" width="20" height="20" v-if="!right">
					<path d="M610.88 512L192 93.12 285.12 0l512 512-512 512L192 930.88z" fill="currentColor" p-id="2587"></path>
				</svg>
				<svg t="1688729212223" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg"
					p-id="1610" width="20" height="20" v-if="right">
					<path d="M378.24 512l418.88 418.88L704 1024 192 512l512-512 93.12 93.12z" fill="currentColor" p-id="1611">
					</path>
				</svg>

			</div>
		</div>
	</div>
</template>
    
<script setup lang='ts'>
import { reactive, ref, onMounted, defineProps, computed } from 'vue'
import { client } from '../http/client'
import { ipcRenderer } from 'electron'
const right = ref(true)//true为右箭头，false为左箭头
// scrollContainer = ref(null);
const checkLight = () => {
	right.value = !right.value
}
// 字体
const fontSize = ref(16);
// 背景
const scrool = ref(null);

const props = defineProps({
	name: {
		type: String,
		default: '暂无',
	},
	local: {
		type: String,
		required: true
	},
	noverid: {
		type: String,
		default: "null"
	}
})
onMounted(() => {
	console.log(props, typeof (props.local))
	if (props.local == '0') {
		client.get(`/api/v1/novel/get/` + props.noverid + `/catalog`, {
			params: {
				token: localStorage.token,
			}
		}).then(
			res => {
				optison.value = res.data.chapters
				console.log(optison.value, '121212')
				changeper(props, 0)
			}
		).catch(
			err => {
				console.log(err)
			}
		)
	} else {
		ipcRenderer.send('Getpassliem', props.name)
	}
})
//章节名
const optison = ref([
	{ value: '第一章：绯红' },
	{ value: '第二章：情况' },
	{ value: '第三章：梅丽莎' },
	{ value: '第四章：占卜' },
	{ value: '第五章：“职业”和住房都是严肃的事情' },
])
// 文件合计
const progressIndex = ref(0)
const percentage = computed(() => {
	return ((progressIndex.value + 1) / optison.value.length * 100).toFixed(2) + "%";
});
const progrdata = ref(null)
function changeper(e, index) {
	progressIndex.value = index
	if (props.local == 0) {
		console.log(optison.value[progressIndex.value])
		progressIndex.value = index
		client.get(`/api/v1/novel/get/` + props.noverid + `/chapter`, {
			params: {
				token: localStorage.token,
				chapterId: optison.value[progressIndex.value].chapterId
			}
		}).then(
			res => {
				// console.log('成功', res.data)
				progrdata.value = processString(res.data)
				// console.log(progrdata.value)
				console.log(scrool.value.scrollTop)
				if (scrool.value) {
					scrool.value.scrollTop = 0;
				}
			}
		).catch(
			err => {
				console.log('失败')
			}
		)
	} else {
		progressIndex.value = index
		getchapter(props.name)
	}
}
//切割文本
function processString(str) {
	const result = [];
	const dotIndex = str.indexOf('\n');

	if (dotIndex !== -1) {
		const substring = str.substring(0, dotIndex);

		result.push(substring);

		const remainingString = str.substring(dotIndex + 1);
		return result.concat(processString(remainingString));
	} else {
		if (str.length > 0) {
			result.push(str);
		}
	}
	return result;
}
function removeExtension(filename) {
	return filename.replace('.txt', '');
}
ipcRenderer.on('Getpassliem', (e, data) => [
	console.log(data),
	optison.value = data.readin,
	getchapter(props.name)
])
function getchapter(bookname) {

	let data = {
		name: bookname,
		chaptername: optison.value[progressIndex.value].name
	}
	// console.log(data.chaptername.name, '212222222222222222222222222')
	ipcRenderer.send('getchapter', data)
}
ipcRenderer.on('chaptersdata', (e, data) => {
	progrdata.value = processString(data)
	// console.log(data)
})
</script>
    
<style scoped lang="scss">
.container {
	height: 100%;
	min-height: auto;
}

.openscr {
	overflow-y: auto;
}

.sidenav {
	position: fixed;
	z-index: 1;
	transition: 0.5s;
	top: 0;
	right: 0;
	height: 100%;
}

p {
	text-indent: 2em;
}

.el-aside {
	height: 100%;
	float: right;
	// width: 286px;
	width: 0;
	background-color: #F2F2F2;
	display: flex;
	flex-direction: column;
	justify-content: space-between;
	transition: 0.5s;
	overflow: hidden;

	.body {
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;

	}


	h2 {
		font-family: '更纱黑体 UI SC';
		font-size: 24px;
	}

	.el-text {
		font-family: '更纱黑体 UI SC';
		font-size: 16px;
	}

	.chapter {
		width: 232px;
		height: 38px;
		line-height: 38px;
		border-radius: 5px;
	}

	.chapter:hover {
		background-color: #D7D7D7;
	}





	.bottom {
		position: absolute;
		background-color: #F2F2F2;
		width: 82%;
		bottom: -4%;
		left: 1vw;
		display: flex;
		align-items: center;
		margin: 20px 10px;

		// 修改input Number样式
		html.dark .el-input-mian {
			background-color: #515151;
		}

		:deep(.el-input-number .el-input__wrapper) {
			// background-color: #D7D7D7;
			border-radius: 100px;

			padding-left: 10px;
			padding-right: 10px;
		}

		:deep(.el-input-number__decrease),
		:deep(.el-input-number__increase) {
			width: 15px;
			height: 30px;
			background: none;
			border: none;
			border-radius: 25%;
			padding: 0 5px
		}



	}


}

html.dark #chapter-ilem:hover {
	background-color: rgb(110, 110, 110) !important;
}


html.dark .bottom {
	background-color: black !important;
}

html.dark #aside {
	background-color: black;
}

.right {
	background-color: #F2F2F2;
	width: 20px;
	height: 50px;
	border: none;
	position: relative;
	z-index: 1;
	// right: 30px;
	margin-top: 340px;
	display: flex;
	align-items: center;
	border-radius: 15px 0 0 15px;
}

html.dark .right {
	background-color: #515151;
}

.right:hover {
	background-color: #b7b7b7;
}

.open {
	width: 286px;
}

.Progress {
	width: 8rem;
	height: 3rem;
	display: flex;
	justify-content: center;
	/* 水平居中 */
	align-items: center;
	/* 垂直居中 */
	position: absolute;
	left: 80vw;
	top: 2vh;
	border-radius: 13px;
	background-color: #ffffff;
}

html.dark .Progress {
	background-color: #121212;
	// background-color: #F2F2F2;
}

.dom {
	// border: 0.5px solid saddlebrown;
	width: 80vw;
	text-indent: 2em;
	font-size: var(--font-size);
}
</style>