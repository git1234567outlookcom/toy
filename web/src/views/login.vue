<style lang="less" scoped>
    @import "./login.less";
</style>
<template>
    <div class="login" @keydown.enter="submit">
        <div class="top">
            <div class="header">
                <a>
                    <img src="../assets/logo.png" class="logo" alt="logo">
                    <span class="title">toy</span>
                </a>
            </div>
            <div class="desc">
            </div>
        </div>
        <div class="main">
            <Form ref="loginForm" :model="user" :rules="rules">
                <FormItem prop="phone" label="账 号">
                    <Input size="large" prefix="ios-person-outline" type="text" v-model="user.phone"
                           placeholder="请输入账号">
                    </Input>
                </FormItem>
                <FormItem prop="password" label="密 码">
                    <Input size="large" prefix="ios-lock-outline" type="password" v-model="user.password"
                           placeholder="请输入密码">
                    </Input>
                </FormItem>
                <FormItem>
                    <Button size="large" @click="submit" type="primary" long>登 录</Button>
                </FormItem>
            </Form>
            <!--			<p class="login-tip">首页-->
            <!--				<a href="/" title="首页">-->
            <!--					<Icon type="ios-send-outline" size="20" />-->
            <!--				</a>-->
            <!--			</p>-->
            <!--			<div class="other-login">-->
            <!--				<span>其他登录方式</span>&nbsp;-->
            <!--				<a>-->
            <!--					<Icon type="logo-github" />-->
            <!--				</a>-->
            <!--			</div>-->
        </div>
        <div class="footer">
            <!--            <div class="links">-->
            <!--                <a>帮助</a>-->
            <!--                <a>隐私</a>-->
            <!--                <a>条款</a>-->
            <!--            </div>-->
            <!--            <div class="copyright">-->
            <!--                Copyright &copy; {{new Date().getFullYear()}}&nbsp;<a target="_blank"-->
            <!--                                                                      href="https://github.com/zxysilent">github.com/zxysilent</a>&nbsp;&nbsp;<a-->
            <!--                    target="_blank" href="https://blog.zxysilent.com">blog.zxysilent.com</a>-->
            <!--            </div>-->
        </div>
    </div>
</template>
<script>
    import {apiLogin} from "@/api/auth";
    import util from "@/utils.js";

    export default {
        data() {
            return {
                user: {phone: "", password: ""},
                rules: {
                    phone: [{required: true, message: "账号不能为空", trigger: "blur"}],
                    password: [{required: true, message: "密码不能为空", trigger: "blur"}]
                }
            };
        },
        methods: {
            submit() {
                let that = this;
                that.$refs.loginForm.validate(valid => {
                    if (valid) {
                        let data = {
                            phone: that.user.phone,
                            password: that.user.password,
                        };
                        apiLogin(data).then(res => {
                            if (res.code == 200) {
                                this.$Message.success({
                                    content: "登陆成功",
                                    onClose: () => {
                                        util.setToken(res.data.token);
                                        that.$router.push({name: "home"});
                                    }
                                });
                            } else {
                                this.$Message.error(res.message);
                            }
                        });
                    }
                });
            }
        }
    };
</script>