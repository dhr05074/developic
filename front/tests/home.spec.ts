import { test, expect, type Page } from '@playwright/test';

test("click Logo",async ({page}) => {
  await page.goto('http://localhost:5173');
  await expect(page).toHaveTitle(/Developic/);
  await page.getByRole('link', { name: 'developic logo image' }).click();
  await expect(page).toHaveURL('http://localhost:5173/');
})

test('Start button change Test', async ({page}) => {
  await page.goto('http://localhost:5173');
  await expect(page).toHaveTitle(/Developic/);

  const button =  await page.getByRole('link', { name: 'Start' })
  await expect(button).toHaveClass(/disabled/);

  await page.locator('div').filter({ hasText: /^Language$/ }).nth(1).click();
  await page.getByText('Javascript').click();
  await expect(button).toHaveClass(/disabled/);

  await page.locator('div').filter({ hasText: /^Javascript$/ }).nth(1).click();
  await page.getByText('Go').click();
  await expect(button).toHaveClass(/disabled/);

  await page.locator('div').filter({ hasText: /^Go$/ }).first().click();
  await page.getByText('Cpp').click();
  await expect(button).toHaveClass(/disabled/);

  
  await page.locator('div').filter({ hasText: /^Difficulty$/ }).nth(1).click();
  await page.getByText('Hard').click();
  await page.locator('div').filter({ hasText: /^Hard$/ }).nth(1).click();
  await page.getByText('Normal').click();
  await page.locator('div').filter({ hasText: /^Normal$/ }).nth(1).click();
  await page.getByText('Easy').click();
  await expect(button).toHaveClass(/normal/);

});
