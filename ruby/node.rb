
class Node
    attr_accessor :left, :right
    attr_reader :value

    def initialize(value)
        @value = value
        @left = nil
        @right = nil
    end
end

def pushNode(node, value)
    if (value > node.value)
        if (node.right)
            pushNode(node.right, value)
        else
            node.right = Node.new(value)
        end
    else
        if (node.left)
            pushNode(node.left, value)
        else
            node.left = Node.new(value)
        end
    end
end

def transverse(tree)
    if tree == nil
        return
    end

    if tree.left
        transverse(tree.left)
    end

    put tree.value

    if tree.right
        transverse(tree.right)
    end
end
