import { test, expect, type Page } from '@playwright/test';

test("profile test",async ({page}) => {
    await page.goto('http://localhost:5173');
    await expect(page).toHaveTitle(/Developic/);
    await page.getByRole('link').nth(1).click();

    await expect(page).toHaveURL(/profile/);

    const scores = await page.getByRole('heading', { name: 'Latest Score' })
    // expect(scores).toHaveLength(9)
    const moreButton = await page.locator('div').filter({ hasId:'profile_More' })
  })
  