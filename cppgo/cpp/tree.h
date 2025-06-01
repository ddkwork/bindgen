#pragma once
#include <vector>
#include <memory>
#include <algorithm>
#include <functional>
#include <string>
#include <sstream>
#include <queue>
#include <stack>
#include <stdexcept>

template<typename T>
struct TreeNode {
    T                       data;    // 节点数据
    int                     depth;   // 节点深度（根节点为0）
    TreeNode               *parent;  // 父节点指针
    std::vector<TreeNode *> children;// 子节点列表

    explicit TreeNode(T val, TreeNode *parent = nullptr) : data(std::move(val)), depth(0), parent(parent) {
        if (parent) { depth = parent->depth + 1; }
    }

    ~TreeNode() {
        for (auto *child: children) { delete child; }
    }

    TreeNode *AddChild(T childData) {
        TreeNode *child = new TreeNode(childData, this);
        children.push_back(child);
        return child;
    }

    void AddChildNode(TreeNode *child) {
        if (child) {
            child->parent = this;
            child->depth = this->depth + 1;
            children.push_back(child);
        }
    }

    TreeNode *RemoveChild(size_t index) {
        if (index >= children.size()) return nullptr;
        TreeNode *removed = children[index];
        children.erase(children.begin() + index);
        return removed;
    }

    void RemoveAllChildren() {
        for (auto *child: children) { delete child; }
        children.clear();
    }

    bool IsLeaf() const { return children.empty(); }

    size_t ChildCount() const { return children.size(); }

    TreeNode *Find(const std::function<bool(const TreeNode *)> &predicate) {
        if (predicate(this)) { return this; }

        for (auto *child: children) {
            TreeNode *result = child->Find(predicate);
            if (result) { return result; }
        }
        return nullptr;
    }

    void WalkDepthFirst(std::function<void(TreeNode *)> visit, bool preOrder = true) {
        if (preOrder) {
            visit(this);// 先访问当前节点
        }

        for (auto *child: children) { child->WalkDepthFirst(visit, preOrder); }

        if (!preOrder) {
            visit(this);// 后访问当前节点
        }
    }

    void WalkBreadthFirst(std::function<void(TreeNode *)> visit) {
        std::queue<TreeNode *> queue;
        queue.push(this);

        while (!queue.empty()) {
            TreeNode *current = queue.front();
            queue.pop();
            visit(current);

            for (auto *child: current->children) { queue.push(child); }
        }
    }

    void SortChildren(std::function<bool(const TreeNode *, const TreeNode *)> compare) {
        std::sort(children.begin(), children.end(), [&](TreeNode *a, TreeNode *b) { return compare(a, b); });
    }

    int Height() const {
        int maxDepth = depth;
        for (auto *child: children) { maxDepth = std::max(maxDepth, child->Height()); }
        return maxDepth - depth;
    }

    virtual std::string ToString() const {
        std::ostringstream oss;
        oss << "TreeNode[depth=" << depth << ", data=" << data;

        if (!children.empty()) { oss << ", children=" << children.size(); }

        if (parent) { oss << ", parent=" << &parent; }

        oss << "]";
        return oss.str();
    }

    std::string FormatTree(const std::string &indent = "") const {
        std::ostringstream oss;
        oss << indent << data.ToString() << "\n";

        std::string childIndent = indent + "  ";
        for (auto *child: children) { oss << child->FormatTree(childIndent); }

        return oss.str();
    }
};

struct PersonData {
    int         id;
    std::string name;
    int         age;

    std::string ToString() const { return "Person{id:" + std::to_string(id) + ", name:" + name + ", age:" + std::to_string(age) + "}"; }

    bool operator<(const PersonData &other) const { return id < other.id; }
};

template<typename T>
struct NTree {
    TreeNode<T> *root;

    NTree() : root(nullptr) {}

    explicit NTree(T rootData) : root(new TreeNode<T>(rootData)) {}

    ~NTree() {
        if (root) { delete root; }
    }

    void SetRoot(TreeNode<T> *newRoot) {
        if (root) { delete root; }
        root = newRoot;
        if (root) {
            root->parent = nullptr;
            root->depth = 0;
        }
    }

    TreeNode<T> *SetRoot(T rootData) {
        SetRoot(new TreeNode<T>(rootData));
        return root;
    }

    bool IsEmpty() const { return root == nullptr; }

    int Height() const {
        if (!root) return 0;
        return root->Height();
    }

    void WalkDepthFirst(std::function<void(TreeNode<T> *)> visit, bool preOrder = true) {
        if (root) { root->WalkDepthFirst(visit, preOrder); }
    }

    void WalkBreadthFirst(std::function<void(TreeNode<T> *)> visit) {
        if (root) { root->WalkBreadthFirst(visit); }
    }

    TreeNode<T> *Find(std::function<bool(TreeNode<T> *)> predicate) {
        if (!root) return nullptr;
        return root->Find(predicate);
    }

    TreeNode<T> *Find(T value) {
        return Find([=](TreeNode<T> *node) { return node->data == value; });
    }

    void DeleteNode(TreeNode<T> *node) {
        if (!node || node == root) return;

        TreeNode<T> *parent = node->parent;
        if (!parent) return;

        auto it = std::remove(parent->children.begin(), parent->children.end(), node);
        parent->children.erase(it, parent->children.end());

        delete node;
    }

    void SortTree(std::function<bool(TreeNode<T> *, TreeNode<T> *)> compare) {
        if (!root) return;

        WalkBreadthFirst([&](TreeNode<T> *node) { node->SortChildren(compare); });
    }

    TreeNode<T> *CloneTree(TreeNode<T> *source) {
        if (!source) return nullptr;

        TreeNode<T> *copy = new TreeNode<T>(source->data);
        copy->depth = source->depth;

        for (auto *child: source->children) {
            TreeNode<T> *childCopy = CloneTree(child);
            if (childCopy) {
                childCopy->parent = copy;
                copy->children.push_back(childCopy);
            }
        }

        return copy;
    }

    std::string ToString() const {
        if (!root) return "NTree[empty]";
        return root->FormatTree();
    }
};

void TestNTree() {
    NTree<std::string> tree;
    tree.SetRoot("Company");

    TreeNode<std::string> *engDept = tree.root->AddChild("Engineering");
    TreeNode<std::string> *salesDept = tree.root->AddChild("Sales");
    TreeNode<std::string> *hrDept = tree.root->AddChild("HR");

    engDept->AddChild("Frontend Team");
    engDept->AddChild("Backend Team");

    TreeNode<std::string> *backendTeam = engDept->children[1];
    backendTeam->AddChild("John (Senior Developer)");
    backendTeam->AddChild("Emma (Junior Developer)");
    backendTeam->AddChild("Alex (DevOps)");

    std::cout << "===== 公司组织结构 =====" << std::endl;
    std::cout << tree.ToString() << std::endl;

    std::cout << "\n===== 广度优先遍历 =====" << std::endl;
    tree.WalkBreadthFirst([](TreeNode<std::string> *node) { std::cout << std::string(node->depth * 2, ' ') << node->data << std::endl; });

    std::cout << "\n===== 深度优先遍历 (前序) =====" << std::endl;
    tree.WalkDepthFirst([](TreeNode<std::string> *node) { std::cout << std::string(node->depth * 2, ' ') << node->data << std::endl; });

    std::cout << "\n===== 深度优先遍历 (后序) =====" << std::endl;
    tree.WalkDepthFirst([](TreeNode<std::string> *node) { std::cout << std::string(node->depth * 2, ' ') << node->data << std::endl; }, false);

    TreeNode<std::string> *node = tree.Find([](TreeNode<std::string> *node) { return node->data.find("Emma") != std::string::npos; });

    if (node) {
        std::cout << "\n找到节点: " << node->data << std::endl;
        std::cout << "节点深度: " << node->depth << std::endl;
        std::cout << "父节点: " << (node->parent ? node->parent->data : "None") << std::endl;
    }

    std::cout << "\n===== 排序前 =====" << std::endl;
    std::cout << backendTeam->FormatTree();

    backendTeam->SortChildren([](TreeNode<std::string> *a, TreeNode<std::string> *b) { return a->data < b->data; });

    std::cout << "\n===== 排序后 =====" << std::endl;
    std::cout << backendTeam->FormatTree();

    struct Department {
        int         id;
        std::string name;
        double      budget;

        std::string ToString() const { return "Dept[" + name + ", id:" + std::to_string(id) + ", $" + std::to_string(budget) + "M]"; }
    };

    NTree<Department>     deptTree;
    TreeNode<Department> *finance = deptTree.SetRoot(Department{1, "Finance", 10.5});
    finance->AddChild(Department{2, "Accounting", 3.2});
    finance->AddChild(Department{3, "Audit", 2.8});

    std::cout << "\n===== 部门树 =====" << std::endl;
    std::cout << deptTree.ToString() << std::endl;

    if (node) {
        std::cout << "\n===== 删除节点: " << node->data << " =====" << std::endl;
        tree.DeleteNode(node);
        std::cout << tree.ToString() << std::endl;
    }

    std::cout << "\n树高: " << tree.Height() << std::endl;

    TreeNode<std::string> *backendClone = tree.CloneTree(backendTeam);
    std::cout << "\n===== 克隆的子树 =====" << std::endl;
    std::cout << backendClone->FormatTree();

    delete backendClone;
}

/*
预期输出：

===== 公司组织结构 =====
Company
  Engineering
    Frontend Team
    Backend Team
      John (Senior Developer)
      Emma (Junior Developer)
      Alex (DevOps)
  Sales
  HR

===== 广度优先遍历 =====
Company
  Engineering
  Sales
  HR
    Frontend Team
    Backend Team
      John (Senior Developer)
      Emma (Junior Developer)
      Alex (DevOps)

===== 深度优先遍历 (前序) =====
Company
  Engineering
    Frontend Team
    Backend Team
      John (Senior Developer)
      Emma (Junior Developer)
      Alex (DevOps)
  Sales
  HR

===== 深度优先遍历 (后序) =====
    Frontend Team
      John (Senior Developer)
      Emma (Junior Developer)
      Alex (DevOps)
    Backend Team
  Engineering
  Sales
  HR
Company

找到节点: Emma (Junior Developer)
节点深度: 3
父节点: Backend Team

===== 排序前 =====
Backend Team
  John (Senior Developer)
  Emma (Junior Developer)
  Alex (DevOps)

===== 排序后 =====
Backend Team
  Alex (DevOps)
  Emma (Junior Developer)
  John (Senior Developer)

===== 部门树 =====
Dept[Finance, id:1, $10.5M]
  Dept[Accounting, id:2, $3.2M]
  Dept[Audit, id:3, $2.8M]

===== 删除节点: Emma (Junior Developer) =====
Company
  Engineering
    Frontend Team
    Backend Team
      John (Senior Developer)
      Alex (DevOps)
  Sales
  HR

树高: 3

===== 克隆的子树 =====
Backend Team
  Alex (DevOps)
  John (Senior Developer)
*/
