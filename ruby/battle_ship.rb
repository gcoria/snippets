def damaged_or_sunk(board, attacks)
  ships = attack_ships(board, attacks)
  h = { 'sunk' => 0, 'damaged' => 0, 'not_touched' => 0 }
  ships.keys.map(&:abs).uniq.each { |s| h[state(ships[s], ships[-s])] += 1 }
  
  h.merge('points' => h['sunk'] + h['damaged'] * 0.5 - h['not_touched'])
end

def attack_ships(board, attacks)
  attacks.each { |x, y| board[-y][x.pred] *= -1 }
  board.flatten.uniq.reject(&:zero?).product([true]).to_h
end

def state(untached, attacked)
  attacked && untached ? 'damaged' : attacked ? 'sunk' : 'not_touched'
end
