import { test, expect } from '@playwright/test';

test('test', async ({ page }) => {
  await page.goto('http://localhost:5173/');
  await page.locator('div').filter({ hasText: /^Language$/ }).nth(1).click();
  await page.getByText('Javascript').click();
  await page.locator('div').filter({ hasText: /^Difficulty$/ }).first().click();
  await page.getByText('Normal').click();
  await page.getByRole('link', { name: 'Start' }).click();
  await page.goto('http://localhost:5173/problem');
  await page.locator('div').filter({ hasText: 'Javascript' }).nth(4).click();
  await page.getByRole('listitem').filter({ hasText: 'Javascript' }).click();
  await page.locator('.cm-content > div:nth-child(15)').click();
  await page.getByText('/** * Definition for singly-linked list. * class ListNode { * val: number * next').fill('\n\n/**\n * Definition for singly-linked list.\n * class ListNode {\n *     val: number\n *     next: ListNode | null\n *     constructor(val?: number, next?: ListNode | null) {\n *         this.val = (val===undefined ? 0 : val)\n *         this.next = (next===undefined ? null : next)\n *     }\n * }\n */\n\n\nfunction addTwoNumbers(l1: ListNode | null, l2: ListNode | null): ListNode | null {\nsdf\n};');
  await page.getByText('/** * Definition for singly-linked list. * class ListNode { * val: number * next').press('CapsLock');
  await page.getByText('/** * Definition for singly-linked list. * class ListNode { * val: number * next').press('CapsLock');
  await page.getByText('/** * Definition for singly-linked list. * class ListNode { * val: number * next').fill('\n\n/**\n * Definition for singly-linked list.\n * class ListNode {\n *     val: number\n *     next: ListNode | null\n *     constructor(val?: number, next?: ListNode | null) {\n *         this.val = (val===undefined ? 0 : val)\n *         this.next = (next===undefined ? null : next)\n *     }\n * }\n */\n\n\nfunction addTwoNumbers(l1: ListNode | null, l2: ListNode | null): ListNode | null {\ncode console.log*""(""????)\n};');
  await page.getByText('/** * Definition for singly-linked list. * class ListNode { * val: number * next').press('Enter');
  await page.getByText('/** * Definition for singly-linked list. * class ListNode { * val: number * next').press('Enter');
  await page.getByText('/** * Definition for singly-linked list. * class ListNode { * val: number * next').fill('\n\n/**\n * Definition for singly-linked list.\n * class ListNode {\n *     val: number\n *     next: ListNode | null\n *     constructor(val?: number, next?: ListNode | null) {\n *         this.val = (val===undefined ? 0 : val)\n *         this.next = (next===undefined ? null : next)\n *     }\n * }\n */\n\n\nfunction addTwoNumbers(l1: ListNode | null, l2: ListNode | null): ListNode | null {\ncode console.log*""(""????\n                   \n                   sdfklsdjfklfs)\n};');
  await page.getByRole('button', { name: 'Reset' }).click();
  await page.locator('.cm-content > div:nth-child(15)').click();
  await page.getByText('* class ListNode {').click();
  await page.locator('.cm-line').first().click();
  await page.getByText('};').click({
    modifiers: ['Shift']
  });
  await page.getByRole('button', { name: 'Reset' }).click();
});