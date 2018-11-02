#include <iostream>
#include <vector>

using namespace std;

class Solution
{
  public:
    void sortColors(vector<int> &nums)
    {
        int a = 0;
        int b = nums.size() - 1;

        for (int i = 0; i <= b; i++)
        {
            if (nums[i] == 1)
            {
                continue;
            }

            if (nums[i] == 0)
            {
                // 0不需要再次判断，直接放到最前端就好
                swap(nums[i], nums[a++]);
            }

            if (nums[i] == 2)
            {
                // 后面的数字换位置后，还需要判断是 0还是1
                swap(nums[i--], nums[b--]);
            }
        }
    }
};

int main(int argc, char const *argv[])
{
    Solution s;

    vector<int> numbers = {1, 2, 0};

    s.sortColors(numbers);

    cout << numbers[0] << "\n";
    cout << numbers[1] << "\n";
    cout << numbers[2] << "\n";

    return 0;
}
