async (page) => {
  await page.goto('http://localhost:8000/login');
  await page.fill('input[type=text]', 'admin');
  await page.fill('input[type=password]', '123456');
  await page.click('button[type=submit]');
  await page.waitForTimeout(4000);

  await page.goto('http://localhost:8000/admin/permission');
  await page.waitForTimeout(4000);

  await page.locator('.ant-table-tbody tr', { hasText: '账号查看' }).first()
    .locator('td:last-child button', { hasText: '绑定接口' }).click();
  await page.waitForTimeout(3000);

  const tree = await page.evaluate(() => {
    const modal = Array.from(document.querySelectorAll('.ant-modal-content')).find((x) => x.querySelector('.ant-tree'));
    if (!modal) return null;
    return Array.from(modal.querySelectorAll('.ant-tree-treenode')).map((n) => {
      const cb = n.querySelector('.ant-tree-checkbox');
      return {
        title: n.querySelector('.ant-tree-title')?.textContent.trim(),
        level: n.querySelectorAll('.ant-tree-indent-unit').length,
        checked: cb?.classList.contains('ant-tree-checkbox-checked') || false,
        indeterminate: cb?.classList.contains('ant-tree-checkbox-indeterminate') || false,
      };
    });
  });

  // 交互验证：勾选一个未勾选接口（账号编辑），再取消，不保存
  let clickResult = null;
  const target = tree.find((t) => t.level === 1 && t.title === '账号编辑（adminUser::edit）');
  if (target) {
    await page.evaluate(() => {
      const modal = Array.from(document.querySelectorAll('.ant-modal-content')).find((x) => x.querySelector('.ant-tree'));
      const nodes = Array.from(modal.querySelectorAll('.ant-tree-treenode'));
      const n = nodes.find((x) => x.querySelector('.ant-tree-title')?.textContent.includes('账号编辑（adminUser::edit）'));
      n?.querySelector('.ant-tree-checkbox')?.click();
    });
    await page.waitForTimeout(800);
    clickResult = await page.evaluate(() => {
      const modal = Array.from(document.querySelectorAll('.ant-modal-content')).find((x) => x.querySelector('.ant-tree'));
      const cb = Array.from(modal.querySelectorAll('.ant-tree-treenode'))
        .find((x) => x.querySelector('.ant-tree-title')?.textContent.includes('账号编辑（adminUser::edit）'))
        ?.querySelector('.ant-tree-checkbox');
      return {
        checked: cb?.classList.contains('ant-tree-checkbox-checked') || false,
        menuIndeterminate: Array.from(modal.querySelectorAll('.ant-tree-treenode'))
          .find((x) => x.querySelector('.ant-tree-title')?.textContent === '账号管理')
          ?.querySelector('.ant-tree-checkbox')?.classList.contains('ant-tree-checkbox-indeterminate') || false,
      };
    });
  }

  return { nodeCount: tree ? tree.length : 0, tree, clickResult };
}
