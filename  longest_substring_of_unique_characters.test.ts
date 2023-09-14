import longestSubstr from "./ longest_substring_of_unique_characters";
import { describe, test, expect  } from "@jest/globals";

describe('longest substring of unique', () => {
    test('', () => {
        expect(longestSubstr("abcabcbb")).toBe(3)
    })
})