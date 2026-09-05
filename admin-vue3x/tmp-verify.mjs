import { chromium } from 'playwright';

const base = 'http://127.0.0.1:8010';
const CHROME = 'C:\\Program Files\\Google\\Chrome\\Application\\chrome.exe';
const results = [];
const check = (name, ok, extra = '') => {
  const line = `${ok ? 'PASS' : 'FAIL'} ${name}${extra ? ' :: ' + extra : ''}`;
  results.push(line);
  console.log(line); // 实时输出，便于脚本中途失败时也能看到已完成项
};

const browser = await chromium.launch({ executablePath: CHROME });
const page = await browser.newPage();

let editCalled = false;
let passwordCalled = false;
page.on('request', (r) => {
  if (r.url().includes('/admin/account/edit')) editCalled = true;
  if (r.url().includes('/admin/account/password')) passwordCalled = true;
});

const errTexts = async () => (await page.locator('.ant-form-item-explain-error').allTextContents()).join(' / ');

// 1. 登录（mock 不校验验证码）
await page.goto(`${base}/login`);
const token = await page.evaluate(async () => {
  const res = await fetch('/api/admin/account/login', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ username: 'admin', password: '123456' }),
  });
  const json = await res.json();
  return json?.data?.token || '';
});
check('mock 登录成功', !!token);
await page.evaluate(
  (t) => localStorage.setItem('token', JSON.stringify({ token: t, expire: Date.now() + 86400000, remember: true })),
  token,
);

// 2. 个人中心
await page.goto(`${base}/account`);
await page.waitForSelector('input[placeholder="请输入昵称"]', { timeout: 20000 });
const SAVE_ACCOUNT = '.account-right button.ant-btn-primary';
await page.waitForSelector(SAVE_ACCOUNT, { timeout: 10000 });

// 场景1：昵称留空直接保存
await page.fill('input[placeholder="请输入昵称"]', '');
await page.click(SAVE_ACCOUNT);
await page.waitForTimeout(800);
check('编辑资料：昵称留空被前端拦下', (await errTexts()).includes('请输入昵称'), `提示="${await errTexts()}"`);
check('编辑资料：昵称留空未请求后端', editCalled === false);

// 场景2：邮箱非法
await page.fill('input[placeholder="请输入昵称"]', 'zhangsan');
await page.fill('input[placeholder="请输入邮箱"]', 'not-an-email');
editCalled = false;
await page.click(SAVE_ACCOUNT);
await page.waitForTimeout(800);
check('编辑资料：非法邮箱被前端拦下', (await errTexts()).includes('邮箱格式不正确'), `提示="${await errTexts()}"`);
check('编辑资料：非法邮箱未请求后端', editCalled === false);

// 场景3：头像留空（本轮改为非必填）应能提交
await page.fill('input[placeholder="请输入邮箱"]', 'a@b.com');
await page.fill('input[placeholder="请输入头像地址"]', '');
editCalled = false;
await page.click(SAVE_ACCOUNT);
await page.waitForTimeout(1500);
check('编辑资料：头像留空仍可提交（非必填）', editCalled === true);

// 4. 修改密码（切到第二个 tab）
const targets = page.getByText('修改密码', { exact: true });
console.log('tab targets:', await targets.count());
await targets.first().click();
await page.waitForTimeout(1200);
console.log('pwd input count after switch:', await page.locator('input[placeholder="请输入原密码"]').count());
await page.waitForSelector('input[placeholder="请输入原密码"]', { timeout: 10000 });
const SAVE_PASSWORD = '.current-password button.ant-btn-primary';
await page.waitForSelector(SAVE_PASSWORD, { timeout: 10000 });
await page.click(SAVE_PASSWORD);
await page.waitForTimeout(800);
check('修改密码：全部留空被前端拦下', (await errTexts()).includes('请输入原密码'), `提示="${await errTexts()}"`);
check('修改密码：全部留空未请求后端', passwordCalled === false);

// 场景5：两次新密码不一致
await page.fill('input[placeholder="请输入原密码"]', 'Old@12345');
await page.fill('input[placeholder="请输入新密码"]', 'New@12345');
await page.fill('input[placeholder="请再次输入新密码"]', 'Other@12345');
passwordCalled = false;
await page.click(SAVE_PASSWORD);
await page.waitForTimeout(800);
check('修改密码：两次密码不一致被拦下', (await errTexts()).includes('两次输入的密码不一致'), `提示="${await errTexts()}"`);
check('修改密码：不一致时未请求后端', passwordCalled === false);

// 场景6：改新密码后错误自动消除
await page.fill('input[placeholder="请输入新密码"]', 'Other@12345');
await page.waitForTimeout(900);
check('修改密码：新密码改一致后错误自动消除', (await errTexts()) === '', `残留提示="${await errTexts()}"`);

// 场景7：合法修改密码应提交
passwordCalled = false;
await page.click(SAVE_PASSWORD);
await page.waitForTimeout(1500);
check('修改密码：合法填写正常提交', passwordCalled === true);

await browser.close();
console.log(results.join('\n'));
console.log(results.every((r) => r.startsWith('PASS')) ? 'ALL CHECKS PASSED' : 'SOME CHECKS FAILED');
