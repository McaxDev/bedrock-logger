<template>
  <a-config-provider :theme="{token: {colorPrimary: '#28abce',borderRadius: 8},}" :locale="locale">
    <RouterView />
  </a-config-provider>
</template>

<script>
import zhCN from 'ant-design-vue/es/locale/zh_CN';
import dayjs from 'dayjs'
import 'dayjs/locale/zh-cn'
dayjs.locale('zh-cn')

export default{
  data(){
    return{
      locale: zhCN,
    }
  }
}
</script>

<style>
/* 滚动条的宽度和轨道样式 */
::-webkit-scrollbar {
  width:4px;
  height: 4px;
}
::-webkit-scrollbar-track {
  background: #ffffff00;
}

/* 滚动条手柄 */
::-webkit-scrollbar-thumb {
  background: #e6e6e6; 
  border-radius: 4px;
}

::-webkit-scrollbar-thumb:hover {
  background: #e6e6e6;
}
</style>
