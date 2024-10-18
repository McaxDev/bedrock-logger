<template>
    <div class="about">
        <a-modal :open="aboutModal" :closeIcon="CloseOutlined" width="100%" wrap-class-name="full-modal" @cancel="closeMoadal">
            <!-- <template #closeIcon>
                <CloseOutlined :style="{fontSize:10}"/>
            </template>  -->
            <template #footer>
                <div style="margin-bottom: 10px;">
                    <a-typography-text code>{{ appInfo.version }}</a-typography-text>
                </div>
                <a-button size="large" @click="closeMoadal">关闭</a-button>
            </template>
            <div class="about-modal">
                <a-typography-title :level="4">关于</a-typography-title>
                <a-typography-paragraph>
                    <a-typography-text>当前为测试版，如有Bug请通过以下邮箱反馈，谢谢！  </a-typography-text>
                    <a-typography-text :copyable="true">nerakolo@outlook.com</a-typography-text>
                </a-typography-paragraph>
                <a-typography-paragraph>
                    <blockquote>
                        <a-typography-text>后端&插件制作：</a-typography-text>
                        <a-typography-link href="https://github.com/bestcb2333" target="_blank">bestcb</a-typography-link>
                    </blockquote>
                    <blockquote>
                        <a-typography-text>前端&设计：</a-typography-text>
                        <a-typography-link href="https://github.com/Nerakolox" target="_blank">Nerakolo</a-typography-link>
                    </blockquote>
                </a-typography-paragraph>
            </div>
        </a-modal>
    </div>
</template>
  
<script>
import { CloseOutlined } from '@ant-design/icons-vue'
import { appInfoStore } from '@/stores/stores'
export default{
    props:['aboutModal'],
    data(){
        return{
            appInfo:appInfoStore()
        }
    },
    methods:{
        handleOk(){
            console.log(this.aboutModal)
        },
        closeMoadal(){
            this.$emit('closeMoadal')
        }
    }
}
</script>

<style scoped>
.about-modal{
    margin-top: 30px;
}
</style>

<style>
.full-modal {
    .ant-modal {
        max-width: 60%;
    }
    .ant-modal-content {
        display: flex;
        flex-direction: column;
        height: 60vh;
        min-height: 460px;
    }
    .ant-modal-body {
        flex: 1;
    }
}
@media (max-width: 992px) {
    .full-modal {
        .ant-modal {
            max-width: 100%;
            top: 0;
            padding-bottom: 0;
            margin: 0;
        }
        .ant-modal-content {
            display: flex;
            flex-direction: column;
            height: calc(100vh);
        }
    }
}
</style>