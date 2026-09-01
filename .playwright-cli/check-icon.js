async (page) => {
  const out = { errors: [] };
  const sleep = (ms) => page.waitForTimeout(ms);
  page.on('pageerror', (e) => out.errors.push('PAGEERROR: ' + String(e).slice(0, 200)));
  page.on('console', (m) => {
    if (m.type() === 'error') out.errors.push('CONSOLE: ' + m.text().slice(0, 200));
  });
  const rowByName = (name) => page.locator('.ant-table-tbody tr:not(.ant-table-measure-row)').filter({ hasText: name }).first();

  await page.goto('http://localhost:8000/login');
  await sleep(1500);
  await page.fill('input[type=text]', 'admin');
  await page.fill('input[type=password]', '123456');
  await page.click('button[type=submit]');
  await sleep(4500);
  await page.goto('http://localhost:8000/admin/menu');
  await sleep(4000);

  try {
    // 新建测试菜单：名称 + 路径 + 图标下拉选择
    await page.locator('button').filter({ hasText: '新建菜单' }).first().click();
    await sleep(1800);
    const inputs = page.locator('.ant-modal-wrap:visible input.ant-input');
    await inputs.nth(0).fill('图标测试菜单');
    await inputs.nth(1).fill('/admin/iconTest');
    await sleep(600);
    out.autoKey = await inputs.nth(2).inputValue();
    const iconItem = page.locator('.ant-modal-wrap:visible .ant-form-item').filter({ hasText: '图标' }).first();
    out.iconIsSelect = await iconItem.locator('.ant-select').count();
    await iconItem.locator('.ant-select').first().click();
    await sleep(1500);
    out.dropdown = await page.evaluate(() => {
      const dd = Array.from(document.querySelectorAll('.ant-select-dropdown')).find((x) => !x.classList.contains('ant-select-dropdown-hidden'));
      if (!dd) return null;
      const items = Array.from(dd.querySelectorAll('.ant-select-item-option'));
      return { rendered: items.length, first: items.slice(0, 3).map((i) => ({ t: i.textContent.trim(), svg: !!i.querySelector('svg') })) };
    });
    await iconItem.locator('input').fill('Setting');
    await sleep(1200);
    out.filtered = await page.evaluate(() => {
      const dd = Array.from(document.querySelectorAll('.ant-select-dropdown')).find((x) => !x.classList.contains('ant-select-dropdown-hidden'));
      return Array.from(dd.querySelectorAll('.ant-select-item-option')).slice(0, 6).map((i) => i.textContent.trim());
    });
    await page.locator('.ant-select-dropdown:not(.ant-select-dropdown-hidden) .ant-select-item-option').first().click();
    await sleep(800);
    out.selected = await iconItem.locator('.ant-select-selection-item').textContent();
    out.selectedSvg = await iconItem.locator('.ant-select-selection-item svg').count();
    await page.locator('.ant-modal-wrap:visible button.ant-btn-primary').click();
    await sleep(2500);
    out.saveMsg = await page.evaluate(() => {
      const n = document.querySelectorAll('.ant-message-notice-content');
      return n.length ? n[n.length - 1].textContent.trim() : '';
    });

    // 编辑回显
    await rowByName('图标测试菜单').locator('button').filter({ hasText: '编辑' }).first().click();
    await sleep(2000);
    const editIcon = page.locator('.ant-modal-wrap:visible .ant-form-item').filter({ hasText: '图标' }).first();
    out.editValue = await editIcon.locator('.ant-select-selection-item').textContent();
    out.editSvg = await editIcon.locator('.ant-select-selection-item svg').count();
    await page.keyboard.press('Escape');
    await sleep(1000);

    // 详情展示
    await rowByName('图标测试菜单').locator('button').filter({ hasText: '详情' }).first().click();
    await sleep(2000);
    out.detailIcon = await page.evaluate(() => {
      const d = Array.from(document.querySelectorAll('.ant-drawer')).find((x) => x.getBoundingClientRect().width > 0);
      const rows = Array.from(d.querySelectorAll('.ant-descriptions-row'));
      const r = rows.find((x) => x.querySelector('.ant-descriptions-item-label')?.textContent?.trim() === '图标');
      const cell = r?.querySelector('.ant-descriptions-item-content');
      return { text: cell?.textContent?.trim(), svg: !!cell?.querySelector('svg') };
    });
    await page.keyboard.press('Escape');
    await sleep(1000);

    // 侧边栏图标（首页）
    out.siderIcon = await page.evaluate(() => {
      const el = document.querySelector('.ant-menu-item a svg, .ant-menu-submenu-title svg');
      return !!el;
    });
  } catch (e) {
    out.error = String(e).slice(0, 300);
  }

  // 清理：禁用 → 删除
  try {
    await rowByName('图标测试菜单').locator('.ant-switch').first().click();
    await sleep(1200);
    await page.locator('.ant-popover:not(.ant-popover-hidden) button.ant-btn-primary').first().click();
    await sleep(2000);
    await rowByName('图标测试菜单').locator('button').filter({ hasText: '删除' }).first().click();
    await sleep(1200);
    await page.locator('.ant-popover:not(.ant-popover-hidden) button.ant-btn-primary').first().click();
    await sleep(2000);
    await page.reload();
    await sleep(3500);
    out.cleanup = await page.evaluate(() => document.body.innerText.includes('图标测试菜单'));
  } catch (e2) {
    out.cleanupError = String(e2).slice(0, 200);
  }
  return out;
}
