async (page) => {
  await page.goto('http://localhost:8000/login');
  await page.fill('input[type=text]', 'admin');
  await page.fill('input[type=password]', '123456');
  await page.click('button[type=submit]');
  await page.waitForTimeout(4000);

  await page.goto('http://localhost:8000/admin/user');
  await page.waitForTimeout(2500);

  // 打开新建账号弹窗
  await page.locator('.ant-btn-primary').last().click();
  await page.waitForSelector('.ant-modal-content', { timeout: 8000 });
  await page.waitForTimeout(1200);

  // 填写已存在的账号 admin
  await page.locator('.ant-modal-content input.ant-input').nth(0).fill('admin');
  await page.locator('.ant-modal-content input[type=password]').first().fill('Test@1234');
  await page.locator('.ant-modal-footer button.ant-btn-primary').click();
  await page.waitForTimeout(2000);

  const msg = await page.evaluate(() => {
    const el = document.querySelector('.ant-message-notice-content, .ant-message-error, .ant-message');
    return el ? el.textContent.trim() : null;
  });
  const modalStillOpen = await page.evaluate(() => {
    const w = document.querySelector('.ant-modal-wrap');
    return w ? w.style.display !== 'none' : false;
  });
  return { msg, modalStillOpen };
}
