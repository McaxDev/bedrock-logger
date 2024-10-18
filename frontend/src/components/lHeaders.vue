<template>
    <div class="l-header">
        <div>
            <!-- <img style="height: 25px;" src="@/assets/axologo.png"> -->
        </div>
        <small :style="{ color:theme.colorPrimary }" style="font-weight: bold;margin-top: 10px;">
            <!-- Axolotland 日志查询系统 -->
            Test Content
        </small>
    </div>
</template>
  
<script>
import { useThemeStore } from '@/stores/stores'
export default{
    data(){
        return{
            theme:useThemeStore().main
        }
    },
    methods:{
        
    }
}
</script>

<style scoped>
.l-header{
    max-width: 100%;
}
</style>