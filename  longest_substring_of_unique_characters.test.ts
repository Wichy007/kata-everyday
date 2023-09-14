import longestSubstr from "./ longest_substring_of_unique_characters";
import { describe, test, expect  } from "@jest/globals";

describe('longest substring of unique', () => {
    test('', () => {
        expect(longestSubstr("abcabcbb")).toBe(3)
        expect(longestSubstr("bbbbb")).toBe(1)
        expect(longestSubstr("pwwkew")).toBe(3)
        expect(longestSubstr("abcdef")).toBe(6)
        expect(longestSubstr("")).toBe(0)
        expect(longestSubstr("au")).toBe(2)
        expect(longestSubstr("dvdf")).toBe(3)
    })
})